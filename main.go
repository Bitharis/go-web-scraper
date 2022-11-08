package main

import (
	"Bitharis/go-web-scraper/internal/miners"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	m := miners.IcaPrototypeMiner{}
	////////////////////////////////////////
	out, _ := os.Create("output.txt")
	defer out.Close()
	resp, err := http.Get(m.GetTargetUrl())
	if err != nil {
		fmt.Println(err.Error())
	}
	var _, _ = io.Copy(out, resp.Body)

	/////////////////////////////////////////
	dataminer := miners.SetupDataminer(&m)
	dataminer.MineData()
	data, error := m.GetMinedData()

	if error != nil {
		fmt.Println(error.Error())
		return
	}

	fmt.Println(data.MinedOnDateTime)
	fmt.Println(data.Source)
	for _, product := range data.Products {
		fmt.Println(product.Name, "Price", product.Price)

	}
}
