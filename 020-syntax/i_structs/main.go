package main

import (
	"encoding/json"
	"fmt"
)

type price struct {
	UnitPrice float32 `json:"betrag-je-stueck"`
	Currency  string  `json:"währung"`
}
type product struct {
	ProductId string `json:"produktId"`
	Name      string `json:"NAME"`
	Price     price  `json:"Preis"`
}

func main() {
	fmt.Println("Structs:")
	p1 := product{
		ProductId: "P-123",
		Name:      "Pizza Salami",
		Price: price{
			UnitPrice: 12.99,
			Currency:  "EUR",
		},
	}
	bytes, _ := json.Marshal(p1)
	fmt.Println(string(bytes))

	bytes, _ = json.MarshalIndent(p1, "", "  ")
	fmt.Println(string(bytes))
}
