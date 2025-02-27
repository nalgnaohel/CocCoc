package business

import (
	"coccoc/internal/item"
	"coccoc/internal/models"
)

// In case we want to change how to get shippingfee with
type ItemBizWithProductType struct {
}

func NewItemBizWithProductType() item.ItemBusiness {
	return &ItemBizWithProductType{}
}

// GetPrice method
func (ib *ItemBizWithProductType) GetPrice(weightCoeff float64, dimensionCoeff float64, item *models.Item) float64 {
	if weightCoeff < 0 || dimensionCoeff < 0 {
		return -1
	}
	if item.AmazonPrice < 0 || item.Weight < 0 || item.ItemDimension.Width < 0 || item.ItemDimension.Height < 0 || item.ItemDimension.Depth < 0 {
		return -1
	}
	return item.AmazonPrice + ib.GetShippingFee(weightCoeff, dimensionCoeff, item)
}

// GetShippingFee method
func (ib *ItemBizWithProductType) GetShippingFee(weightCoeff float64, dimensionCoeff float64, item *models.Item) float64 {
	return max(ib.GetFeeByWeight(weightCoeff, item), ib.GetFeeByDimension(dimensionCoeff, item), ib.GetFeeByProductType(item))
}

func (ItemBizWithProductType) GetFeeByProductType(item *models.Item) float64 {
	return 100 //this is just an example, not real implementation
}

// GetFeeByWeight method
func (ib *ItemBizWithProductType) GetFeeByWeight(weightCoeff float64, item *models.Item) float64 {
	return weightCoeff * item.Weight
}

// GetFeeByDimension method
func (ib *ItemBizWithProductType) GetFeeByDimension(dimensionCoeff float64, item *models.Item) float64 {
	return dimensionCoeff * (item.ItemDimension.Width * item.ItemDimension.Height * item.ItemDimension.Depth)
}
