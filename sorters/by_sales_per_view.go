package sorters

import (
	"product-catalog-sorter/models"
	"sort"
)

// SortBySalesPerView sorts by sales_count / views_count ratio
type SortBySalesPerView struct{}

func (s SortBySalesPerView) Sort(products []models.Product) {
	sort.Slice(products, func(i, j int) bool {
		ratioI := float64(products[i].SalesCount) / float64(products[i].ViewsCount)
		ratioJ := float64(products[j].SalesCount) / float64(products[j].ViewsCount)
		return ratioI > ratioJ
	})
}
