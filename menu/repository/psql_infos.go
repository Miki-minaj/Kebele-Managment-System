package repository

import (
	"database/sql"
	"errors"
	"fmt"

	//"github.com/betsegawlemma/restaurant/entity"
	"github.com/miki-minaj/Kebele-Managment-System/entity"
)

// CategoryRepositoryImpl implements the menu.CategoryRepository interface
type CategoryRepositoryImpl struct {
	conn *sql.DB
}

// NewCategoryRepositoryImpl will create an object of PsqlCategoryRepository
func NewCategoryRepositoryImpl(Conn *sql.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{conn: Conn}
}
func (cri *CategoryRepositoryImpl) StoreCategory(c *entity.Category) (*entity.Category, error) {
	id := c.ID
	name := c.Name

	fmt.Println(id)
	fmt.Println(name)
	sqlStatement := `
	INSERT INTO infos (id,name)
	VALUES (` + id + `,'` + name + `')`

	_, err := cri.conn.Query(sqlStatement)
	if err != nil {
		fmt.Println("this  is")
		return nil, errors.New("Insertion has failed")
	}

	return nil, nil
}

func (cri *CategoryRepositoryImpl) Categories() ([]entity.Category, error) {

	rows, err := cri.conn.Query("SELECT * FROM infos;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	ctgs := []entity.Category{}

	for rows.Next() {
		category := entity.Category{}
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, category)
	}

	return ctgs, nil
}
