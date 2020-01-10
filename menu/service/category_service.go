package service

import (
	"fmt"

	"github.com/miki-minaj/Kebele-Managment-System/entity"
	"github.com/miki-minaj/Kebele-Managment-System/menu"
)

// CategoryService implements menu.CategoryService interface
type CategoryService struct {
	categoryRepo menu.CategoryRepository
}

// NewCategoryService will create new CategoryService object
func NewCategoryService(CatRepo menu.CategoryRepository) menu.CategoryService {
	return &CategoryService{categoryRepo: CatRepo}
}

// StoreCategory persists new category information
func (cs *CategoryService) StoreCategory(category *entity.Category) (*entity.Category, error) {

	cat, errs := cs.categoryRepo.StoreCategory(category)
	if errs != nil {
		fmt.Println("and now")
		return nil, errs
	}

	return cat, nil
}

func (cs *CategoryService) Categories() ([]entity.Category, error) {

	categories, errs := cs.categoryRepo.Categories()

	if errs != nil {
		return nil, errs
	}

	return categories, nil
}
