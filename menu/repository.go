package menu

import "github.com/miki-minaj/Kebele-Managment-System/entity"

// CategoryRepository specifies food menu category database operations
type CategoryRepository interface {
	StoreCategory(category *entity.Category) (*entity.Category, error)
	Categories() ([]entity.Category, error)
}
