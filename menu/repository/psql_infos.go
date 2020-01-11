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
	mother := c.Mothername
	age := c.AGE
	occup := c.Occupation
	relegion := c.Relegion
	nation := c.Nationality
	phone := c.Phonenumber
	emer := c.Emergencyname
	emerph := c.Emergencyphone
	sex := "male"
	fmt.Println(id)
	fmt.Println(name)
	fmt.Println(mother)
	fmt.Println(age)
	fmt.Println(occup)
	fmt.Println(relegion)
	fmt.Println(nation)
	fmt.Println(phone)
	fmt.Println(emer)
	fmt.Println(emerph)
	sqlStatement := `
	INSERT INTO infos (id,name,mothername,sex,age,religion,occupations,phonenum,nationality,emrgecyname,emergecyphone)
	VALUES (` + id + `,'` + name + `','` + mother + `','` + sex + `',` + age + `,'` + relegion + `','` + occup + `',` + phone + `,'` + nation + `','` + emer + `','` + emerph + `')`

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
		err = rows.Scan(&category.ID, &category.Name, &category.Mothername, &category.Sex, &category.AGE, &category.Relegion, &category.Occupation, &category.Phonenumber, &category.Nationality, &category.Emergencyname, &category.Emergencyphone)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, category)
	}

	return ctgs, nil
}
