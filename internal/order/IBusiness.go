package order

import (
	"coccoc/internal/models"
)

type OrderBusiness interface {
	GetGrossPrice(weightCoeff float64, dimensionCoeff float64, order *models.Order) float64
}
