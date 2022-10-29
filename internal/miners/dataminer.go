package miners

import (
	"fmt"

	"github.com/gocolly/colly"
)

func requestHandler(request *colly.Request) {
	request.Headers.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	request.Headers.Add("Accept-Encoding", "gzip, deflate, br")
	request.Headers.Add("Accept-Language", "en-US,en;q=0.5")
	request.Headers.Add("Connection", "keep-alive")
	request.Headers.Add("Sec-Fetch-Dest", "document")
	request.Headers.Add("Sec-Fetch-Mode", "navigate")
	request.Headers.Add("Sec-Fetch-Site", "none")
	request.Headers.Add("Sec-Fetch-User", "?1")
	request.Headers.Add("Upgrade-Insecure-Requests", "1")
	request.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:106.0) Gecko/20100101 Firefox/106.0")
	fmt.Println("Visiting", request.URL)
}
func responseHandler(response *colly.Response) { fmt.Println("Got", response.Request.URL) }

func errorHandler(r *colly.Response, err error) { fmt.Print("Error: ", r.Request.URL, err) }

type DataMiner struct {
	rootUrl                string
	selector               string
	onSelectorFoundHandler colly.HTMLCallback
	collector              colly.Collector
}

func SetupDataminer(m IMiner) *DataMiner {
	collector := colly.NewCollector()
	collector.OnRequest(requestHandler)
	collector.OnResponse(responseHandler)
	collector.OnError(errorHandler)
	return &DataMiner{
		rootUrl:                m.GetTargetUrl(),
		selector:               m.GetTargetRootHtmlElement(),
		onSelectorFoundHandler: m.OnTargetFound,
		collector:              *collector,
	}
}

func (dataminer *DataMiner) MineData() {
	dataminer.collector.OnHTML(dataminer.selector, dataminer.onSelectorFoundHandler)
	dataminer.collector.Visit(dataminer.rootUrl)
}
