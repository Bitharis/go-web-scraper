package miners

import (
	"github.com/gocolly/colly"
)

type IMiner interface {
	OnTargetFound(e *colly.HTMLElement)
	GetTargetUrl() string
	GetTargetRootHtmlElement() string
	GetMinedData() (*Data, error)
	GetRequestHeaders() (headers map[string]string)
}
