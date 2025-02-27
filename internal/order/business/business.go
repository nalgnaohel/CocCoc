package business

import (
	"coccoc/internal/item"
	"coccoc/internal/models"
	"coccoc/internal/order"
)

type OrderBiz struct {
	itemBiz item.ItemBusiness
}

func NewOrderBiz(ib item.ItemBusiness) order.OrderBusiness {
	return &OrderBiz{
		itemBiz: ib,
	}
}

// GetGrossPrice method
func (ob *OrderBiz) GetGrossPrice(weightCoeff float64, dimensionCoeff float64, order *models.Order) float64 {
	var grossPrice float64
	if order == nil || len(order.Items) == 0 {
		return 0
	}
	for _, item := range order.Items {
		grossPrice += ob.itemBiz.GetPrice(weightCoeff, dimensionCoeff, &item)
	}
	if grossPrice < 0 {
		return 0
	}
	return grossPrice
}
