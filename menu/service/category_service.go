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
func (cs *CategoryService) Category(id int) (entity.Category, error) {

	c, err := cs.categoryRepo.Category(id)

	if err != nil {
		return c, err
	}
	return c, nil
}

// DeleteCategory delete a category by its id
func (cs *CategoryService) DeleteCategory(id int) error {

	err := cs.categoryRepo.DeleteCategory(id)
	if err != nil {
		return err
	}
	return nil
}
