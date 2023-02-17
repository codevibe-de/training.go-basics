package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type Ingredient struct {
	name       string
	prepTime   int
	isPrepared bool
}

func (ing *Ingredient) prepare() {
	fmt.Println("Preparing ", ing.name)
	time.Sleep(time.Duration(ing.prepTime) * time.Second)
	ing.isPrepared = true
}

type Pizza struct {
	ingredients []*Ingredient
	bakingTime  int
}

type Oven struct {
	prepTime   int
	isPrepared bool
}

func (o *Oven) prepare() {
	fmt.Println("Heating up oven...")
	time.Sleep(time.Duration(o.prepTime) * time.Second)
	o.isPrepared = true
}

func (o *Oven) bake(pizza Pizza) error {
	if !o.isPrepared {
		return errors.New("oven is not prepared")
	}
	for _, ing := range pizza.ingredients {
		if !ing.isPrepared {
			return fmt.Errorf("ingredient %q is not prepared", ing.name)
		}
	}
	time.Sleep(time.Duration(pizza.bakingTime) * time.Second)
	fmt.Println("Pizza is done!")
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
		ingredients: []*Ingredient{dough1, tomatoSauce1, gratedCheese1},
	}
	pizzaSalamiSpecial = Pizza{
		ingredients: []*Ingredient{dough2, tomatoSauce2, gratedCheese2, salami2, onions2},
	}
)

var oven = Oven{8, false}

func main() {
	// we want to measure time until both pizza are done
	startTime := time.Now()

	// prep oven
	oven.prepare()

	// prep ingredients
	for _, p := range []Pizza{pizzaMargarita, pizzaSalamiSpecial} {
		for _, ing := range p.ingredients {
			ing.prepare()
		}
	}

	// lets bake some pizza
	err := oven.bake(pizzaMargarita)
	handleError("baking pizza", err)
	err = oven.bake(pizzaSalamiSpecial)
	handleError("baking pizza", err)

	// measure
	fmt.Printf("All pizza done after %d seconds", int32(time.Since(startTime).Seconds()))
}

func handleError(context string, err error) {
	if err != nil {
		fmt.Printf("Got error %s: %s", context, err.Error())
		os.Exit(2)
	}
}
