package interfaces

import (
	"github.com/amir-mln/go-projects/micro-cafe/product-management/core/product"
)

type ProductRepository interface {
	// FetchByID(int) (*product.Entity, error)
	FetchAll() ([]*product.Entity, error)
	// Insert(*product.Entity) error
	// Update(*product.Entity) error
	// Delete(int) error
}
