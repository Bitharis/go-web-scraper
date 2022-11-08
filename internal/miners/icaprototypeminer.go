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

var icatargetUrl string = "https://www.ica.se/handla/kategori/kott--fagel---fisk-id_1/?s=ica-kvantum-malmborgs-clemenstorget-id_02771"
var icatargetRootElement = ".main-container"
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
	fmt.Println(e.DOM.Html())
	var name = "test"
	var price = e.ChildText("p")

	normalizedPriceString := normalizepricestring(price)
	p, err := strconv.ParseFloat(normalizedPriceString, 64)

	if err != nil {
		fmt.Println(err.Error())
	}

	var product = model.Product{
		Name:  name,
		Price: p,
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
	headers["Host"] = "www.ica.se"
	headers["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"
	headers["Accept-Encoding"] = "gzip, deflate, br"
	headers["Accept-Language"] = "en-US,en;q=0.5"
	headers["Connection"] = "keep-alive"
	headers["Sec-Fetch-Dest"] = "document"
	headers["Sec-Fetch-Mode"] = "navigate"
	headers["Sec-Fetch-Site"] = "none"
	headers["Sec-Fetch-User"] = "?1"
	headers["Upgrade-Insecure-Requests"] = "1"
	headers["Cookie"] = "userPrefLanguage=sv_SE; hpc=523d9836288b9dfa6eff415927079d23_1667755905; TS0154f30f=01f0ddaba39bc3089dcc53a8376ac29c2f0e79ca503f4991b6ecdf7c29c1e5f8a9ca268c017838ee9b0c9591e907a9bd3ffb3225ef; OptanonConsent=isGpcEnabled=0&datestamp=Sun+Nov+06+2022+18%3A32%3A15+GMT%2B0100+(Central+European+Standard+Time)&version=6.33.0&isIABGlobal=false&hosts=&consentId=72f20165-2a10-43fb-aaf0-d6fd297f04ec&interactionCount=1&landingPath=NotLandingPage&groups=C0001%3A1%2CC0003%3A0%2CC0002%3A0%2CC0004%3A0&iType=2&geolocation=%3B&AwaitingReconsent=false; OptanonAlertBoxClosed=2022-11-06T17:31:59.668Z; JSESSIONID=GWxN_7nAYcNsPdpxwh0mBpsedzaLDnPQg4ZgQDw8qVPi4PhmzCg_!2056063056; EPiStateMarker=true"
	return headers
}
