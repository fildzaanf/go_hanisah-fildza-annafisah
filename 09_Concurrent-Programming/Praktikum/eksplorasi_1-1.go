package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Product struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func productData(urlAPI string, wg *sync.WaitGroup, productChannel chan<- Product) {
	defer wg.Done()

	respon, err := http.Get(urlAPI)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer respon.Body.Close()

	var products []Product
	err = json.NewDecoder(respon.Body).Decode(&products)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, product := range products {
		productChannel <- product
	}
}

func main() {
	urlAPI := "https://fakestoreapi.com/products"
	productChannel := make(chan Product)
	var wg sync.WaitGroup

	wg.Add(1)
	go productData(urlAPI, &wg, productChannel)

	go func() {
		defer close(productChannel)
		wg.Wait()
	}()

	fmt.Println("Product data")

	for product := range productChannel {
		fmt.Println("===")
		fmt.Printf("Title: %s\nPrice: %.2f\nCategory: %s\n", product.Title, product.Price, product.Category)
		fmt.Println("===")
	}
}
