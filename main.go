package main

import (
	"Bitharis/go-web-scraper/internal/miners"
	"fmt"
)

func main() {
	m := miners.IcaPrototypeMiner{}
	dataminer := miners.SetupDataminer(&m)
	dataminer.MineData()
	data, error := m.GetMinedData()

	if error != nil {
		fmt.Println(error.Error())
	}

	fmt.Println(data.MinedOnDateTime)
	fmt.Println(data.Source)
	for _, product := range data.Products {
		fmt.Println(product.Name, "Price", product.Price)

	}
}
