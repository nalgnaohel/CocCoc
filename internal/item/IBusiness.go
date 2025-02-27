package item

import (
	"coccoc/internal/models"
)

// ItemBusiness interface that includes GetPrice, GetShippingFee, and GetShippingFeeExtension interfaces
type ItemBusiness interface {
	GetPrice(weightCoeff float64, dimensionCoeff float64, item *models.Item) float64
	GetShippingFee(weightCoeff float64, dimensionCoeff float64, item *models.Item) float64
	GetFeeByWeight(weightCoeff float64, item *models.Item) float64
	GetFeeByDimension(dimensionCoeff float64, item *models.Item) float64
}
