package menu

type MenuItem struct {
	ID      int `json:"id"`
	MenuRef int `json:"-"`
	// Menu        Menu    `json:"menu"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}
