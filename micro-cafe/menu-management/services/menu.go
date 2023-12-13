package services

import (
	"github.com/amir-mln/go-projects/micro-cafe/menu-management/core/menu"
	"github.com/amir-mln/go-projects/micro-cafe/menu-management/core/menu/interfaces"
)

type MenuBusiness struct {
	menuRepo interfaces.MenuRepository
}

func (pb *MenuBusiness) GetAllMenus() ([]*menu.Menu, error) {
	return pb.menuRepo.FetchAll()
}

func NewMenuBusiness(menuRepo interfaces.MenuRepository) *MenuBusiness {
	return &MenuBusiness{menuRepo}
}
