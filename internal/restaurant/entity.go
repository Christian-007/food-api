package restaurant

import (
	"food/internal/menu"
)

type Restaurant struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
	Menus []menu.Menu `json:"menus"`
}
