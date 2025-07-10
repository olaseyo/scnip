package sorters

import (
	"product-catalog-sorter/models"
	"sort"
)

// SortByPrice sorts products by ascending price
type SortByPrice struct{}

func (s SortByPrice) Sort(products []models.Product) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
}
