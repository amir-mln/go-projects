package repositories

import (
	"encoding/json"
	"os"
	"path"

	"github.com/amir-mln/go-projects/micro-cafe/menu-management/core/menu"
)

type MenuRepository struct{}

func (pr *MenuRepository) FetchAll() ([]*menu.Menu, error) {
	wd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	dataPathSegments := []string{
		wd,
		"infrastructure",
		"repositories",
		"menu.json",
	}

	d, err := os.ReadFile(path.Join(dataPathSegments...))

	if err != nil {
		return nil, err
	}

	menus := &[]*menu.Menu{}

	if err := json.Unmarshal(d, menus); err != nil {
		return nil, err
	}

	return *menus, nil
}

func NewMenuRepository() *MenuRepository {
	return &MenuRepository{}
}
