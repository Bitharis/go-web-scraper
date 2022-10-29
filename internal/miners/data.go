package miners

import (
	"Bitharis/go-web-scraper/internal/model"
	"time"
)

type Data struct {
	MinedOnDateTime time.Time

	Source string

	Products []model.Product
}
