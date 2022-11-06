package miners

import (
	"fmt"

	"github.com/gocolly/colly"
)

var customMiner IMiner

func requestHandler(request *colly.Request) {

	// Set the headers for each miner
	for key, value := range customMiner.GetRequestHeaders() {
		request.Headers.Add(key, value)
	}

	// Set Browser
	request.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:106.0) Gecko/20100101 Firefox/106.0")
}
func responseHandler(response *colly.Response) { fmt.Println("Got", response.Request.URL) }

func errorHandler(r *colly.Response, err error) { fmt.Print("Error: ", r.Request.URL, err) }

type DataMiner struct {
	rootUrl                string
	selector               string
	onSelectorFoundHandler colly.HTMLCallback
	collector              colly.Collector
}

func SetupDataminer(miner IMiner) *DataMiner {
	customMiner = miner
	collector := colly.NewCollector()
	collector.OnRequest(requestHandler)
	collector.OnResponse(responseHandler)
	collector.OnError(errorHandler)
	
	return &DataMiner{
		rootUrl:                customMiner.GetTargetUrl(),
		selector:               customMiner.GetTargetRootHtmlElement(),
		onSelectorFoundHandler: customMiner.OnTargetFound,
		collector:              *collector,
	}
}

func (dataminer *DataMiner) MineData() {
	fmt.Println("Visiting", dataminer.rootUrl)
	dataminer.collector.OnHTML(dataminer.selector, dataminer.onSelectorFoundHandler)
	dataminer.collector.Visit(dataminer.rootUrl)
}
