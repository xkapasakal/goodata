package main

import (
	"fmt"
	"github.com/xkapasakal/go4OData/generated"
	"encoding/json"
)

type DType struct {
	Next string `json:"__next"`
	Results []generated.Customer
}

type Response struct {
	D DType
}


func main() {

	var jsonBlob = generated.Customers()
	var response Response
	err := json.Unmarshal(jsonBlob, &response)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("customers: %+v\n", response.D)
	fmt.Printf("count customers: %+v\n", len(response.D.Results))

	var animalJsonBlob = []byte(`{"d": { "__next": "Kapas"}}`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals Response
	err = json.Unmarshal(animalJsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n ", animals)
}