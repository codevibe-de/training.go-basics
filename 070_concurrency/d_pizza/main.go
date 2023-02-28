package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type Preparer interface {
	fmt.Stringer
	prepare()
}

type Ingredient struct {
	name       string
	prepTime   int
	isPrepared bool
}

func (ing *Ingredient) prepare() {
	fmt.Println("Preparing", ing.name)
	time.Sleep(time.Duration(ing.prepTime) * time.Second)
	ing.isPrepared = true
}

func (i *Ingredient) String() string {
	return i.name
}

type Pizza struct {
	name        string
	ingredients []*Ingredient
	bakingTime  int
}

func (p Pizza) readyToBake() bool {
	for _, ing := range p.ingredients {
		if !(*ing).isPrepared {
			return false
		}
	}
	return true
}

type Oven struct {
	prepTime   int
	isPrepared bool
}

func (o *Oven) prepare() {
	fmt.Println("Oven heating up...")
	time.Sleep(time.Duration(o.prepTime) * time.Second)
	o.isPrepared = true
}

func (o *Oven) String() string {
	return "Oven"
}

func (o *Oven) bake(pizza Pizza) error {
	fmt.Println("Trying to bake pizza", pizza.name)
	if !o.isPrepared {
		return errors.New("oven is not prepared")
	}
	for _, ing := range pizza.ingredients {
		if !ing.isPrepared {
			return fmt.Errorf("ingredient %q is not prepared", ing.name)
		}
	}
	time.Sleep(time.Duration(pizza.bakingTime) * time.Second)
	fmt.Printf("Pizza %s is done!\n", pizza.name)
	return nil
}

var (
	dough1        = &Ingredient{"Dough", 4, false}
	tomatoSauce1  = &Ingredient{"Tomato Sauce", 1, false}
	gratedCheese1 = &Ingredient{"Grated Cheese", 2, false}
	dough2        = &Ingredient{"Dough", 4, false}
	tomatoSauce2  = &Ingredient{"Tomato Sauce", 1, false}
	gratedCheese2 = &Ingredient{"Grated Cheese", 2, false}
	salami2       = &Ingredient{"Salami", 1, false}
	onions2       = &Ingredient{"Onions", 3, false}
)

var (
	pizzaMargarita = Pizza{
		"Margarita",
		[]*Ingredient{dough1, tomatoSauce1, gratedCheese1},
		5,
	}
	pizzaSalamiSpecial = Pizza{
		"Salami Speciale",
		[]*Ingredient{dough2, tomatoSauce2, gratedCheese2, salami2, onions2},
		6,
	}
)

var oven = Oven{8, false}

func main() {
	// we want to measure time until both pizza are done
	startTime := time.Now()

	// launch workers
	log.Println("Pizzeria starting up!")
	pizzaInCh, pizzaOutCh := startPizzeria(&oven, 5)

	// bake them
	// ps of course we could use a for loop here if we had more pizzas
	pizzaInCh <- pizzaMargarita
	pizzaInCh <- pizzaSalamiSpecial

	// wait for 2 pizzas to complete -- this relies on the pizza-out channel to be closed eventually
	remaining := 2
	for p := range pizzaOutCh {
		log.Printf("Pizza %s is done and can be served\n", p.name)
		remaining--
		if remaining == 0 {
			break
		}
	}

	// measure
	log.Printf("All pizza done after %d seconds\n", int32(time.Since(startTime).Seconds()))
}

func startPizzeria(oven *Oven, workerCount int) (chan Pizza, chan Pizza) {
	// create channels
	prepInCh := make(chan Preparer, 100)
	prepOutCh := make(chan Preparer)
	pizzaInCh := make(chan Pizza, 10)
	pizzaOutCh := make(chan Pizza)

	// launch some workers
	for n := 0; n < workerCount; n++ {
		go worker(n+1, prepInCh, pizzaInCh, prepOutCh, pizzaOutCh, oven)
	}

	// launch supervisor checking for prepared ingredients to trigger pizza baking
	go supervisor(prepInCh, prepOutCh, pizzaInCh, oven)

	// get oven going
	prepInCh <- oven

	// done
	return pizzaInCh, pizzaOutCh
}

func worker(
	id int,
	prepInCh chan Preparer,
	pizzaInCh <-chan Pizza,
	prepOutCh chan<- Preparer,
	pizzaOutCh chan<- Pizza,
	oven *Oven,
) {
	for {
		log.Printf("W%d: waiting for work\n", id)
		select {
		case p, ok := <-prepInCh:
			if !ok {
				log.Printf("W%d: detected closed pizza input channel\n", id)
				prepInCh = nil
			} else {
				log.Printf("W%d: got %s to prepare\n", id, p.String())
				p.prepare()
				prepOutCh <- p
			}
		case p, ok := <-pizzaInCh:
			if !ok {
				log.Printf("W%d: detected closed ingredients input channel\n", id)
				pizzaInCh = nil
			} else {
				// we received a pizza - this can happen in TWO cases:
				// 1) pizza has been ordered, ingredients need to be prepared
				// 2) all ingredients have been prepared and pizza can be baked
				if !p.readyToBake() {
					log.Printf("W%d: adding ingredients for pizza %s to work list\n", id, p.name)
					for _, ing := range p.ingredients {
						prepInCh <- ing
					}
				} else {
					log.Printf("W%d: got pizza %s to bake\n", id, p.name)
					err := oven.bake(p)
					if err != nil {
						log.Panicf("W%d: error baking pizza %s: %s\n", id, p.name, err.Error())
					} else {
						pizzaOutCh <- p
					}
				}
			}
		}
		if prepInCh == nil && pizzaInCh == nil {
			log.Printf("W%d: all channels closed -- worker's going home\n", id)
			return
		}
	}
}

func supervisor(prepInCh chan Preparer, prepOutCh chan Preparer, pizzaInCh chan<- Pizza, oven *Oven) {
	unbakedPizzas := []Pizza{pizzaMargarita, pizzaSalamiSpecial}
	for obj := range prepOutCh {
		switch obj.(type) {
		case Preparer:
			for i, pizza := range unbakedPizzas {
				if pizza.readyToBake() && oven.isPrepared {
					log.Printf("SV: Pizza %s ready to bake, handing over to workers\n", pizza.name)
					fmt.Println(unbakedPizzas)
					unbakedPizzas[i] = unbakedPizzas[len(unbakedPizzas)-1]
					unbakedPizzas = unbakedPizzas[:len(unbakedPizzas)-1]
					pizzaInCh <- pizza
				}
			}
		}
		if len(unbakedPizzas) == 0 {
			close(prepInCh)
			close(prepOutCh)
			close(pizzaInCh)
		}
	}
}

func handleError(context string, err error) {
	if err != nil {
		log.Printf("Got error %s: %s", context, err.Error())
		os.Exit(2)
	}
}
