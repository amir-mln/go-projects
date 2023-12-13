package menu

type Menu struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	MenuItems   []MenuItem `json:"menu-items"`
	CreatedOn   string     `json:"-"`
	UpdatedOn   string     `json:"-"`
	DeletedOn   string     `json:"-"`
}
