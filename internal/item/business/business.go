package business

import (
	"coccoc/internal/item"
	"coccoc/internal/models"
)

type ItemBiz struct {
}

func NewItemBiz() item.ItemBusiness {
	return &ItemBiz{}
}

// GetPrice method
func (ib *ItemBiz) GetPrice(weightCoeff float64, dimensionCoeff float64, item *models.Item) float64 {
	if weightCoeff < 0 || dimensionCoeff < 0 {
		return -1
	}
	if item.AmazonPrice < 0 || item.Weight <= 0 || item.ItemDimension.Width <= 0 || item.ItemDimension.Height <= 0 || item.ItemDimension.Depth <= 0 {
		return -1
	}
	return item.AmazonPrice + ib.GetShippingFee(weightCoeff, dimensionCoeff, item)
}

// GetShippingFee method
func (ib *ItemBiz) GetShippingFee(weightCoeff float64, dimensionCoeff float64, item *models.Item) float64 {
	if weightCoeff < 0 || dimensionCoeff < 0 {
		return -1
	}
	if item.AmazonPrice < 0 || item.Weight <= 0 || item.ItemDimension.Width <= 0 || item.ItemDimension.Height <= 0 || item.ItemDimension.Depth <= 0 {
		return -1
	}
	return max(ib.GetFeeByWeight(weightCoeff, item), ib.GetFeeByDimension(dimensionCoeff, item))
}

// GetFeeByWeight method
func (ib *ItemBiz) GetFeeByWeight(weightCoeff float64, item *models.Item) float64 {
	if weightCoeff < 0 {
		return -1
	}
	if item.AmazonPrice < 0 || item.Weight <= 0 || item.ItemDimension.Width <= 0 || item.ItemDimension.Height <= 0 || item.ItemDimension.Depth <= 0 {
		return -1
	}
	return weightCoeff * item.Weight
}

// GetFeeByDimension method
func (ib *ItemBiz) GetFeeByDimension(dimensionCoeff float64, item *models.Item) float64 {
	if dimensionCoeff < 0 {
		return -1
	}
	if item.AmazonPrice < 0 || item.Weight <= 0 || item.ItemDimension.Width <= 0 || item.ItemDimension.Height <= 0 || item.ItemDimension.Depth <= 0 {
		return -1
	}
	return dimensionCoeff * (item.ItemDimension.Width * item.ItemDimension.Height * item.ItemDimension.Depth)
}
