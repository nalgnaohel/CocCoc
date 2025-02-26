package models

type Item struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	AmazonPrice float64 `json:"price"`
	Weight      float64 `json:"weight"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	Depth       float64 `json:"depth"`
}

// use variadic function so that we ca pass any number of parameters
// without changing the code
func (i *Item) GetShippingFee(feeParams ...float64) float64 {
	result := -1.0
	for _, param := range feeParams {
		result = max(result, param)
	}
	return result
}

// Get item price
func (i *Item) GetPrice(feeParams ...float64) float64 {
	return i.AmazonPrice + i.GetShippingFee(feeParams...)
}

// Get fee by weight
func (i *Item) GetFeeByWeight(weightCoeff float64) float64 {
	return i.Weight * weightCoeff
}

// Get fee by dimension
func (i *Item) GetFeeByDimension(dimensionCoeff float64) float64 {
	return i.Width * i.Height * i.Depth * dimensionCoeff
}
