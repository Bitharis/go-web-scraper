package miners

import (
	"Bitharis/go-web-scraper/internal/model"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

var targetUrl string = "https://www.netonnet.se/art/dator-surfplatta/laptop"
var targetRootElement = ".trackingProduct.hidden"
var laptops []model.Product

func normalizePriceString(priceWithComma string) string {
	return strings.Replace(priceWithComma, ",", ".", 1)
}

type NetOnNetLaptopMiner struct {
}

func (*NetOnNetLaptopMiner) GetTargetUrl() string {
	return targetUrl
}

func (*NetOnNetLaptopMiner) GetTargetRootHtmlElement() string {
	return targetRootElement
}

func (*NetOnNetLaptopMiner) OnTargetFound(e *colly.HTMLElement) {
	var productIdElement = e.DOM.Find("[name=\"ProductId\"]")
	var productNameElement = e.DOM.Find("[name=\"ProductName\"]")
	var productPriceElement = e.DOM.Find("[name=\"ProductPrice\"]")
	var productBrandElement = e.DOM.Find("[name=\"ProductBrand\"]")

	var id string
	var name string
	var price float64
	var brand string

	for _, attribute := range productIdElement.Nodes[0].Attr {
		if attribute.Key == "value" {
			id = attribute.Val
		}
	}

	for _, attribute := range productNameElement.Nodes[0].Attr {
		if attribute.Key == "value" {
			name = attribute.Val
		}
	}

	for _, attribute := range productPriceElement.Nodes[0].Attr {
		if attribute.Key == "value" {
			normalizedPriceString := normalizePriceString(attribute.Val)
			p, err := strconv.ParseFloat(normalizedPriceString, 64)

			if err != nil {
				fmt.Println(err.Error())
			}

			price = p
		}
	}

	for _, attribute := range productBrandElement.Nodes[0].Attr {
		if attribute.Key == "value" {
			brand = attribute.Val
		}
	}

	var laptop = model.Product{
		TrackingId: id,
		Name:       name,
		Price:      price,
		Brand:      brand,
	}

	laptops = append(laptops, laptop)
}

func (*NetOnNetLaptopMiner) GetMinedData() (*Data, error) {
	if len(laptops) == 0 {
		err := errors.New("no data mined")
		return nil, err
	}

	var minedData = Data{
		MinedOnDateTime: time.Now().UTC(),
		Source:          "NetOnNet",
		Products:        laptops,
	}

	return &minedData, nil
}
