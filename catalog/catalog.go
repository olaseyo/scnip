package catalog

import (
	"product-catalog-sorter/models"
	"product-catalog-sorter/sorters"
)

type Catalog struct {
	Products []models.Product
	Sorter   sorters.Sorter
}

func (c *Catalog) Sort() {
	c.Sorter.Sort(c.Products)
}
