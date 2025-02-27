package business

import (
	"coccoc/internal/item/business"
	"coccoc/internal/models"
	"testing"
)

func TestGetGrossPrice(t *testing.T) {
	tests := []struct {
		name               string
		items              []models.Item
		weightCoeff        float64
		dimensionCoeff     float64
		expectedGrossPrice float64
	}{
		{
			name: "Test Case 1",
			items: []models.Item{
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
			},
			weightCoeff:        1.0,
			dimensionCoeff:     1.0,
			expectedGrossPrice: 100.0 + 24.0 + 200.0 + 60.0,
		},
		{
			name: "Test Case 2",
			items: []models.Item{
				{
					ID:          3,
					Name:        "Test Item 3",
					AmazonPrice: 300.0,
					Weight:      30.0,
					ItemDimension: models.ItemDimension{
						Width:  4.0,
						Height: 5.0,
						Depth:  6.0,
					},
				},
				{
					ID:          4,
					Name:        "Test Item 4",
					AmazonPrice: 400.0,
					Weight:      40.0,
					ItemDimension: models.ItemDimension{
						Width:  5.0,
						Height: 6.0,
						Depth:  7.0,
					},
				},
			},
			weightCoeff:        1.5,
			dimensionCoeff:     2.0,
			expectedGrossPrice: 300.0 + 240.0 + 400.0 + 420.0,
		},
	}

	mockItemBiz := &business.ItemBiz{}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := models.NewOrder(1, tt.items)
			orderBiz := NewOrderBiz(mockItemBiz)
			grossPrice := orderBiz.GetGrossPrice(tt.weightCoeff, tt.dimensionCoeff, order)
			if grossPrice != tt.expectedGrossPrice {
				t.Errorf("test case %d: expected %v, got %v", i, tt.expectedGrossPrice, grossPrice)
			}
		})
	}
}
