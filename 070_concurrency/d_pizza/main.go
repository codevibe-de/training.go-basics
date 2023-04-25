package main

import (
	"errors"
	"fmt"
	"golang.org/x/exp/slices"
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
	log.Printf("%s is being prepared\n", ing.name)
	time.Sleep(time.Duration(ing.prepTime) * time.Second)
	ing.isPrepared = true
}

func (ing *Ingredient) String() string {
	return ing.name
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
	log.Println("Oven is heating up...")
	time.Sleep(time.Duration(o.prepTime) * time.Second)
	o.isPrepared = true
}

func (o *Oven) String() string {
	return "Oven"
}

func (o *Oven) bake(pizza *Pizza) error {
	log.Printf("Trying to bake Pizza %s\n", pizza.name)
	if !o.isPrepared {
		return errors.New("oven is not prepared")
	}
	for _, ing := range pizza.ingredients {
		if !ing.isPrepared {
			return fmt.Errorf("ingredient %q is not prepared", ing.name)
		}
	}
	time.Sleep(time.Duration(pizza.bakingTime) * time.Second)
	log.Printf("Pizza %s is done!\n", pizza.name)
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
	pizzaMargarita = &Pizza{
		"Margarita",
		[]*Ingredient{dough1, tomatoSauce1, gratedCheese1},
		5,
	}
	pizzaSalamiSpecial = &Pizza{
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
	pizzaOrderCh, pizzaDeliverCh := startPizzeria(&oven, 3)

	// bake them
	// ps of course we could use a for loop here if we had more pizzas
	pizzaOrderCh <- pizzaMargarita
	pizzaOrderCh <- pizzaSalamiSpecial
	close(pizzaOrderCh)

	// wait for pizzas to complete -- this relies on the pizza-out channel to be closed eventually
	for p := range pizzaDeliverCh {
		log.Printf("Pizza %s is done and can be served\n", p.name)
	}

	// measure
	log.Printf("All pizza done after %d seconds\n", int32(time.Since(startTime).Seconds()))
}

func startPizzeria(oven *Oven, workerCount int) (chan *Pizza, chan *Pizza) {
	// create channels
	pizzaOrderCh := make(chan *Pizza, 10)
	pizzaDeliverCh := make(chan *Pizza)
	prepInCh := make(chan Preparer, 100)
	prepOutCh := make(chan Preparer)
	pizzaInCh := make(chan *Pizza, 10)
	pizzaOutCh := make(chan *Pizza)

	// launch some workers
	for n := 0; n < workerCount; n++ {
		go worker(n+1, prepInCh, pizzaInCh, prepOutCh, pizzaOutCh, oven)
	}

	// launch supervisor checking for prepared ingredients to trigger pizza baking
	go supervisor(pizzaOrderCh, pizzaDeliverCh, prepInCh, prepOutCh, pizzaInCh, pizzaOutCh, oven)

	// get oven going
	prepInCh <- oven

	// done
	return pizzaOrderCh, pizzaDeliverCh
}

func worker(
	id int,
	prepInCh <-chan Preparer,
	pizzaInCh <-chan *Pizza,
	prepOutCh chan<- Preparer,
	pizzaOutCh chan<- *Pizza,
	oven *Oven,
) {
	for {
		log.Printf("Worker #%d: waiting for work\n", id)
		select {
		case p, ok := <-prepInCh:
			if !ok {
				log.Printf("Worker #%d: detected closed preparer input channel\n", id)
				prepInCh = nil
			} else {
				log.Printf("Worker #%d: got %s to prepare\n", id, p.String())
				p.prepare()
				prepOutCh <- p
			}
		case p, ok := <-pizzaInCh:
			if !ok {
				log.Printf("Worker #%d: detected closed pizza input channel\n", id)
				pizzaInCh = nil
			} else {
				log.Printf("Worker #%d: got pizza %s to bake\n", id, p.name)
				err := oven.bake(p)
				if err != nil {
					log.Panicf("Worker #%d: error baking pizza %s: %s\n", id, p.name, err.Error())
				} else {
					pizzaOutCh <- p
				}
			}
		}
		if prepInCh == nil && pizzaInCh == nil {
			log.Printf("Worker #%d: all channels closed -- worker's going home\n", id)
			return
		}
	}
}

func supervisor(pizzaOrderCh <-chan *Pizza, pizzaDeliverChan chan<- *Pizza,
	prepInCh chan<- Preparer, prepOutCh <-chan Preparer,
	pizzaInCh chan<- *Pizza, pizzaOutCh <-chan *Pizza,
	oven *Oven) {
	unbakedPizzas := make([]*Pizza, 0)
	undeliveredPizzas := make([]*Pizza, 0)
	for {
		select {
		case pizza, ok := <-pizzaOrderCh:
			if ok {
				log.Printf("SV: Pizza %s has been ordered, handing ingredients over to workers\n", pizza.name)
				unbakedPizzas = append(unbakedPizzas, pizza)
				undeliveredPizzas = append(undeliveredPizzas, pizza)
				for _, ing := range pizza.ingredients {
					prepInCh <- ing
				}
			}
		case ing, ok := <-prepOutCh:
			if ok {
				log.Printf("SV: Got notified that %s has been prepared\n", ing.String())
				if oven.isPrepared {
					for i := 0; i < len(unbakedPizzas); i++ {
						pizza := unbakedPizzas[i]
						if pizza.readyToBake() {
							log.Printf("SV: Pizza %s is ready to bake, handing over to workers\n", pizza.name)
							unbakedPizzas = slices.Delete(unbakedPizzas, i, i+1)
							pizzaInCh <- pizza
							if len(unbakedPizzas) == 0 {
								log.Println("SV: All pizzas are baked, closing some channels")
								close(prepInCh)
								close(pizzaInCh)
							}
							i--
						}
					}
				}
			}
		case pizza, ok := <-pizzaOutCh:
			if ok {
				log.Printf("SV: Got notified that Pizza %s has been baked, informing customer\n", pizza.name)
				pizzaDeliverChan <- pizza
				i := slices.IndexFunc(undeliveredPizzas, func(p *Pizza) bool { return p == pizza })
				undeliveredPizzas = slices.Delete(undeliveredPizzas, i, i+1)
			}
		}
		if len(undeliveredPizzas) == 0 {
			log.Println("SV: All pizzas are delivered, closing some channels and going home")
			close(pizzaDeliverChan)
			return
		}
	}
}

func handleError(context string, err error) {
	if err != nil {
		log.Printf("Got error %s: %s", context, err.Error())
		os.Exit(2)
	}
}
