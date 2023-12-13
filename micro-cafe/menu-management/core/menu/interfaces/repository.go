package interfaces

import (
	"github.com/amir-mln/go-projects/micro-cafe/menu-management/core/menu"
)

type MenuRepository interface {
	// FetchByID(int) (*menu.Menu, error)
	FetchAll() ([]*menu.Menu, error)
	// Insert(*menu.Menu) error
	// Update(*menu.Menu) error
	// Delete(int) error
}
