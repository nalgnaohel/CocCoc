package models

type Order struct {
	ID         int     `json:"id"`
	Items      []Item  `json:"items"`
	GrossPrice float64 `json:"gross_price"`
}

// Get total price of an order
func (o *Order) GetTotalPrice(weightCoeff, dimensionCoeff float64) float64 {
	totalPrice := 0.0
	for _, item := range o.Items {
		totalPrice += item.GetPrice(item.GetFeeByWeight(weightCoeff), item.GetFeeByDimension(dimensionCoeff))
	}
	return totalPrice
}
