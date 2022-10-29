package main

import (
	"Bitharis/go-web-scraper/internal/miners"
	"fmt"
)

func main() {
	m := miners.NetOnNetLaptopMiner{}
	dataminer := miners.SetupDataminer(&m)
	dataminer.MineData()
	data, error := m.GetMinedData()

	if error != nil {
		fmt.Println(error.Error())
	}

	fmt.Println(data.MinedOnDateTime)
	fmt.Println(data.Source)
	for _, product := range data.Products {
		fmt.Println("ID", product.TrackingId, "Brand", product.Brand, "Name", product.Name, "Price", product.Price)

	}
}
