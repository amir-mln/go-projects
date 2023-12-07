package services

import (
	"github.com/amir-mln/go-projects/micro-cafe/product-management/core/product"
	"github.com/amir-mln/go-projects/micro-cafe/product-management/core/product/interfaces"
)

type ProductBusiness struct {
	productRepo interfaces.ProductRepository
}

func (pb *ProductBusiness) GetAllProducts() ([]*product.Entity, error) {
	return pb.productRepo.FetchAll()
}

func NewProductBusiness(productRepo interfaces.ProductRepository) *ProductBusiness {
	return &ProductBusiness{productRepo}
}
