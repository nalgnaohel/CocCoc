package main

import (
	itemBusiness "coccoc/internal/item/business"
	"coccoc/internal/models"
	orderBusiness "coccoc/internal/order/business"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Input weight coefficient: ")
	var weightCoeff float64
	fmt.Scanln(&weightCoeff)
	fmt.Println("Input dimension coefficient: ")
	var dimensionCoeff float64
	fmt.Scanln(&dimensionCoeff)

	items := []models.Item{
		{
			ID:          1,
			Name:        "Test Item 1",
			AmazonPrice: 100.0,
			Weight:      10.0,
			ItemDimension: models.ItemDimension{
				Width:  2.0,
				Height: 3.0,
				Depth:  4.0,
			},
		},
		{
			ID:          2,
			Name:        "Test Item 2",
			AmazonPrice: 200.0,
			Weight:      20.0,
			ItemDimension: models.ItemDimension{
				Width:  3.0,
				Height: 4.0,
				Depth:  5.0,
			},
		},
	}
	order := models.NewOrder(1, items)
	itemBiz := itemBusiness.NewItemBiz()
	orderBiz := orderBusiness.NewOrderBiz(itemBiz)

	grossPrice := orderBiz.GetGrossPrice(weightCoeff, dimensionCoeff, order)
	fmt.Println("Gross price: ", grossPrice)
}
