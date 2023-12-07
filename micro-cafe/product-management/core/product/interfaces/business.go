package interfaces

import "github.com/amir-mln/go-projects/micro-cafe/product-management/core/product"

type ProductBusiness interface {
	GetAllProducts() ([]*product.Entity, error)
}
