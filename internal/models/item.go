package models

type Item struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	AmazonPrice   float64 `json:"price"`
	Weight        float64 `json:"weight"`
	ItemDimension `json:"dimension"`
}

type ItemDimension struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Depth  float64 `json:"depth"`
}

func NewItem(id int, name string, price, weight float64,
	dimension ItemDimension, weightCoeff float64, dimensionCoeff float64) *Item {
	return &Item{
		ID:          id,
		Name:        name,
		AmazonPrice: price,
		Weight:      weight,
		ItemDimension: ItemDimension{
			Width:  dimension.Width,
			Height: dimension.Height,
			Depth:  dimension.Depth,
		},
	}
}
