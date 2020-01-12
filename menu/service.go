package menu

import "github.com/miki-minaj/Kebele-Managment-System/entity"

// CategoryService specifies food menu category services
type InfoService interface {
	StoreCategory(category *entity.Category) (*entity.Category, error)
	Categories() ([]entity.Category, error)
	DeleteCategory(id int) error
	Category(id int) (entity.Category, error)
}
