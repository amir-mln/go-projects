package interfaces

import "github.com/amir-mln/go-projects/micro-cafe/menu-management/core/menu"

type MenuBusiness interface {
	GetAllMenus() ([]*menu.Menu, error)
}
