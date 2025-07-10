package main

import (
	"fmt"
	"product-catalog-sorter/catalog"
	"product-catalog-sorter/models"
	"product-catalog-sorter/sorters"
)

func main() {
	products := []models.Product{
		{ID: 1, Name: "Alabaster Table", Price: 12.99, Created: "2019-01-04", SalesCount: 32, ViewsCount: 730},
		{ID: 2, Name: "Zebra Table", Price: 44.49, Created: "2012-01-04", SalesCount: 301, ViewsCount: 3279},
		{ID: 3, Name: "Coffee Table", Price: 10.00, Created: "2014-05-28", SalesCount: 1048, ViewsCount: 20123},
	}

	// Sort by price
	priceCatalog := catalog.Catalog{Products: products, Sorter: sorters.SortByPrice{}}
	priceCatalog.Sort()
	fmt.Println("Sorted by Price:")
	for _, p := range priceCatalog.Products {
		fmt.Println(p.Name, p.Price)
	}

	// Sort by sales per view
	ratioCatalog := catalog.Catalog{Products: products, Sorter: sorters.SortBySalesPerView{}}
	ratioCatalog.Sort()
	fmt.Println("\nSorted by Sales per View:")
	for _, p := range ratioCatalog.Products {
		ratio := float64(p.SalesCount) / float64(p.ViewsCount)
		fmt.Printf("%s (%.4f)\n", p.Name, ratio)
	}
}
