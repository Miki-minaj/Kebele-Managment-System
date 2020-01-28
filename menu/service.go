package menu

import "github.com/miki-minaj/Kebele-Managment-System/entity"

// CategoryService specifies food menu category services
type CategoryService interface {
	StoreCategory(category *entity.Category) (*entity.Category, error)
	Categories() ([]entity.Category, error)
	DeleteCategory(id int) error
	Category(name string) ([]entity.Category, error)
}
