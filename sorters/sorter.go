package sorters

import "product-catalog-sorter/models"

// Sorter is the strategy interface for sorting
type Sorter interface {
	Sort(products []models.Product)
}
