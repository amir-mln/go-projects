package repositories

import (
	"encoding/json"
	"os"
	"path"

	"github.com/amir-mln/go-projects/micro-cafe/product-management/core/product"
)

type ProductRepository struct{}

func (pr *ProductRepository) FetchAll() ([]*product.Entity, error) {
	wd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	dataPathSegments := []string{
		wd,
		"product-management",
		"drivers",
		"repositories",
		"product.json",
	}

	d, err := os.ReadFile(path.Join(dataPathSegments...))

	if err != nil {
		return nil, err
	}

	products := &[]*product.Entity{}

	if err := json.Unmarshal(d, products); err != nil {
		return nil, err
	}

	return *products, nil
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}
