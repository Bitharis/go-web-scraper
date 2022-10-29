package model

/*Represent a product that was mined from a webpage*/
type Product struct {
	// Id retrieved from the target web site
	TrackingId string

	Name string

	Description string

	Price float64

	Brand string
}
