package app

import (
	"github.com/amir-mln/go-projects/micro-cafe/product-management/core/product/interfaces"
	"github.com/amir-mln/go-projects/micro-cafe/product-management/services"
)

type ApplicationService struct {
	PBiz *services.ProductBusiness
}

func NewApplicationService(pr interfaces.ProductRepository) *ApplicationService {
	return &ApplicationService{PBiz: services.NewProductBusiness(pr)}
}
