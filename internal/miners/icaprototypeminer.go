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

var icatargetUrl string = "https://www.ica.se/handla/kategori/mejeri---ost-id_256/?s=ica-supermarket-linero-torg-id_15172"
var icatargetRootElement = "#application-main"
var products []model.Product

func normalizepricestring(priceWithComma string) string {
	return strings.Replace(priceWithComma, ",", ".", 1)
}

type IcaPrototypeMiner struct {
}

func (*IcaPrototypeMiner) GetTargetUrl() string {
	return icatargetUrl
}

func (*IcaPrototypeMiner) GetTargetRootHtmlElement() string {
	return icatargetRootElement
}

func (*IcaPrototypeMiner) OnTargetFound(e *colly.HTMLElement) {

	var element = e.DOM.Find("ul > li")
	fmt.Println(element)
	var name = "test"
	var price = e.ChildText("p")

	normalizedPriceString := normalizepricestring(price)
			p, err := strconv.ParseFloat(normalizedPriceString, 64)

			if err != nil {
				fmt.Println(err.Error())
			}

	var product = model.Product{
		Name:       name,
		Price:      p,
	}

	products = append(products, product)
}

func (*IcaPrototypeMiner) GetMinedData() (*Data, error) {
	if len(products) == 0 {
		err := errors.New("no data mined")
		return nil, err
	}

	var minedData = Data{
		MinedOnDateTime: time.Now().UTC(),
		Source:          "ICA",
		Products:        products,
	}

	return &minedData, nil
}

func (*IcaPrototypeMiner) GetRequestHeaders() (h map[string]string) {
	headers := make(map[string]string)
	headers["Accept"]="text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"
	headers["Accept-Encoding"]="gzip, deflate, br"
	headers["Accept-Language"]="en-US,en;q=0.5"
	headers["Connection"]="keep-alive"
	headers["Sec-Fetch-Dest"]="document"
	headers["Sec-Fetch-Mode"]="navigate"
	headers["Sec-Fetch-Site"]="none"
	headers["Sec-Fetch-User"]="?1"
	headers["Upgrade-Insecure-Requests"]="1"
	return headers
}
