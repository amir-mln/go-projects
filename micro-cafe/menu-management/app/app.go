package app

import (
	"github.com/amir-mln/go-projects/micro-cafe/menu-management/core/menu/interfaces"
	"github.com/amir-mln/go-projects/micro-cafe/menu-management/services"
)

type ApplicationService struct {
	MenuBusiness *services.MenuBusiness
}

func NewApplicationService(mr interfaces.MenuRepository) *ApplicationService {
	return &ApplicationService{MenuBusiness: services.NewMenuBusiness(mr)}
}
