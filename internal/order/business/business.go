package business

import (
	"coccoc/internal/item"
	"coccoc/internal/models"
	"coccoc/internal/order"
)

type OrderBiz struct {
	itemBiz item.ItemBusiness
}

func NewOrderBiz(itemBiz item.ItemBusiness) order.OrderBusiness {
	return &OrderBiz{
		itemBiz: itemBiz,
	}
}

// GetGrossPrice method
func (ob *OrderBiz) GetGrossPrice(weightCoeff float64, dimensionCoeff float64, order *models.Order) float64 {
	var grossPrice float64
	for _, item := range order.Items {
		grossPrice += ob.itemBiz.GetPrice(weightCoeff, dimensionCoeff, &item)
	}
	return grossPrice
}
