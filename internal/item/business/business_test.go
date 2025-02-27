package business

import (
	"coccoc/internal/models"
	"testing"
)

func TestGetPrice(t *testing.T) {
	tests := []struct {
		name           string
		item           models.Item
		weightCoeff    float64
		dimensionCoeff float64
	}{
		{
			name: "Test Case 1",
			item: models.Item{
				ID:          1,
				Name:        "Test Item 1",
				AmazonPrice: -1.25,
				Weight:      10.0,
				ItemDimension: models.ItemDimension{
					Width:  2.0,
					Height: 3.0,
					Depth:  4.0,
				},
			},
			weightCoeff:    1.0,
			dimensionCoeff: 1.0,
		},
		{
			name: "Test Case 2",
			item: models.Item{
				ID:          2,
				Name:        "Test Item 2",
				AmazonPrice: 200.0,
				Weight:      0.0,
				ItemDimension: models.ItemDimension{
					Width:  3.0,
					Height: 4.0,
					Depth:  5.0,
				},
			},
			weightCoeff:    1.5,
			dimensionCoeff: 2.0,
		},
		{
			name: "Test Case 3",
			item: models.Item{
				ID:          3,
				Name:        "Test Item 3",
				AmazonPrice: 200.0,
				Weight:      5.0,
				ItemDimension: models.ItemDimension{
					Width:  -1.0,
					Height: 4.0,
					Depth:  5.0,
				},
			},
			weightCoeff:    1.5,
			dimensionCoeff: 2.0,
		},
		{
			name: "Test Case 4",
			item: models.Item{
				ID:          4,
				Name:        "Test Item 4",
				AmazonPrice: 200.0,
				Weight:      0.0,
				ItemDimension: models.ItemDimension{
					Width:  3.0,
					Height: -100.0,
					Depth:  5.0,
				},
			},
			weightCoeff:    1.5,
			dimensionCoeff: 2.0,
		},
		{
			name: "Test Case 5",
			item: models.Item{
				ID:          5,
				Name:        "Test Item 5",
				AmazonPrice: 200.0,
				Weight:      0.0,
				ItemDimension: models.ItemDimension{
					Width:  3.0,
					Height: 4.0,
					Depth:  -5.0,
				},
			},
			weightCoeff:    1.5,
			dimensionCoeff: 2.0,
		},
		{
			name: "Test Case 6",
			item: models.Item{
				ID:          6,
				Name:        "Test Item 6",
				AmazonPrice: 200.0,
				Weight:      0.0,
				ItemDimension: models.ItemDimension{
					Width:  3.0,
					Height: 4.0,
					Depth:  5.0,
				},
			},
			weightCoeff:    0,
			dimensionCoeff: 2.0,
		},
		{
			name: "Test Case 7",
			item: models.Item{
				ID:          7,
				Name:        "Test Item 7",
				AmazonPrice: 200.0,
				Weight:      0.0,
				ItemDimension: models.ItemDimension{
					Width:  3.0,
					Height: 4.0,
					Depth:  5.0,
				},
			},
			weightCoeff:    2.5,
			dimensionCoeff: -2.0,
		},
		{
			name: "Test Case 8",
			item: models.Item{
				ID:          8,
				Name:        "Test Item 8",
				AmazonPrice: 200.0,
				Weight:      5.0,
				ItemDimension: models.ItemDimension{
					Width:  3.0,
					Height: 4.0,
					Depth:  5.0,
				},
			},
			weightCoeff:    20.0,
			dimensionCoeff: 2.0,
		},
	}

	itemBiz := NewItemBiz()

	for ind, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			price := itemBiz.GetPrice(tt.weightCoeff, tt.dimensionCoeff, &tt.item)
			if ind >= 0 && ind < 7 {
				if price != -1 {
					t.Errorf("expected -1, got %v", price)
				}
			} else {
				if price != 320.0 {
					t.Errorf("expected %v, got %v", tt.item.AmazonPrice+itemBiz.GetShippingFee(tt.weightCoeff, tt.dimensionCoeff, &tt.item), price)
				}
			}
		})
	}
}

func TestGetShippingFee(t *testing.T) {
	tests := []struct {
		name                string
		item                models.Item
		weightCoeff         float64
		dimensionCoeff      float64
		expectedShippingFee float64
	}{
		{
			name: "Test Case 1",
			item: models.Item{
				ID:          1,
				Name:        "Test Item 1",
				AmazonPrice: -1.25,
				Weight:      10.0,
				ItemDimension: models.ItemDimension{
					Width:  2.0,
					Height: 3.0,
					Depth:  4.0,
				},
			},
			weightCoeff:         1.0,
			dimensionCoeff:      1.0,
			expectedShippingFee: -1,
		},
		{
			name: "Test Case 2",
			item: models.Item{
				ID:          2,
				Name:        "Test Item 2",
				AmazonPrice: 200.0,
				Weight:      0.0,
				ItemDimension: models.ItemDimension{
					Width:  3.0,
					Height: 4.0,
					Depth:  5.0,
				},
			},
			weightCoeff:         1.5,
			dimensionCoeff:      2.0,
			expectedShippingFee: -1,
		},
		{
			name: "Test Case 3",
			item: models.Item{
				ID:          3,
				Name:        "Test Item 3",
				AmazonPrice: 200.0,
				Weight:      5.0,
				ItemDimension: models.ItemDimension{
					Width:  -1.0,
					Height: 4.0,
					Depth:  5.0,
				},
			},
			weightCoeff:         1.5,
			dimensionCoeff:      2.0,
			expectedShippingFee: -1,
		},
		{
			name: "Test Case 4",
			item: models.Item{
				ID:          4,
				Name:        "Test Item 4",
				AmazonPrice: 200.0,
				Weight:      0.0,
				ItemDimension: models.ItemDimension{
					Width:  3.0,
					Height: -100.0,
					Depth:  5.0,
				},
			},
			weightCoeff:         1.5,
			dimensionCoeff:      2.0,
			expectedShippingFee: -1,
		},
		{
			name: "Test Case 5",
			item: models.Item{
				ID:          5,
				Name:        "Test Item 5",
				AmazonPrice: 200.0,
				Weight:      0.0,
				ItemDimension: models.ItemDimension{
					Width:  3.0,
					Height: 4.0,
					Depth:  -5.0,
				},
			},
			weightCoeff:         1.5,
			dimensionCoeff:      2.0,
			expectedShippingFee: -1,
		},
		{
			name: "Test Case 6",
			item: models.Item{
				ID:          6,
				Name:        "Test Item 6",
				AmazonPrice: 200.0,
				Weight:      0.0,
				ItemDimension: models.ItemDimension{
					Width:  3.0,
					Height: 4.0,
					Depth:  5.0,
				},
			},
			weightCoeff:         0,
			dimensionCoeff:      2.0,
			expectedShippingFee: -1,
		},
		{
			name: "Test Case 7",
			item: models.Item{
				ID:          7,
				Name:        "Test Item 7",
				AmazonPrice: 200.0,
				Weight:      0.0,
				ItemDimension: models.ItemDimension{
					Width:  3.0,
					Height: 4.0,
					Depth:  5.0,
				},
			},
			weightCoeff:         2.5,
			dimensionCoeff:      -2.0,
			expectedShippingFee: -1,
		},
		{
			name: "Test Case 8",
			item: models.Item{
				ID:          8,
				Name:        "Test Item 8",
				AmazonPrice: 100.0,
				Weight:      10.0,
				ItemDimension: models.ItemDimension{
					Width:  2.0,
					Height: 3.0,
					Depth:  4.0,
				},
			},
			weightCoeff:         1.0,
			dimensionCoeff:      1.0,
			expectedShippingFee: 24.0,
		},
		{
			name: "Test Case 9",
			item: models.Item{
				ID:          9,
				Name:        "Test Item 9",
				AmazonPrice: 200.0,
				Weight:      20.0,
				ItemDimension: models.ItemDimension{
					Width:  3.0,
					Height: 4.0,
					Depth:  5.0,
				},
			},
			weightCoeff:         1.5,
			dimensionCoeff:      2.0,
			expectedShippingFee: 120.0,
		},
	}

	itemBiz := NewItemBiz()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shippingFee := itemBiz.GetShippingFee(tt.weightCoeff, tt.dimensionCoeff, &tt.item)
			if shippingFee != tt.expectedShippingFee {
				t.Errorf("expected %v, got %v", tt.expectedShippingFee, shippingFee)
			}
		})
	}
}
