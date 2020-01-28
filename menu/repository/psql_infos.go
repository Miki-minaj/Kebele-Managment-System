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
	im := c.Image
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
	INSERT INTO infos (id,name,mothername,sex,age,religion,occupations,phonenum,nationality,emrgecyname,emergecyphone,image)
	VALUES (` + id + `,'` + name + `','` + mother + `','` + sex + `',` + age + `,'` + relegion + `','` + occup + `',` + phone + `,'` + nation + `','` + emer + `',` + emerph + `,'` + im + `')`

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
		err = rows.Scan(&category.ID, &category.Name, &category.Mothername, &category.Sex, &category.AGE, &category.Relegion, &category.Occupation, &category.Phonenumber, &category.Nationality, &category.Emergencyname, &category.Emergencyphone, &category.Image)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, category)
	}

	return ctgs, nil
}

// Category returns a category with a given id
func (cri *CategoryRepositoryImpl) Category(name string) ([]entity.Category, error) {
	fmt.Println("um trying")
	satement := `SELECT * FROM infos WHERE name ='` + name + `'`
	row, _ := cri.conn.Query(satement)

	ctgs := []entity.Category{}

	for row.Next() {
		category := entity.Category{}
		err := row.Scan(&category.ID, &category.Name, &category.Mothername, &category.Sex, &category.AGE, &category.Relegion, &category.Occupation, &category.Phonenumber, &category.Nationality, &category.Emergencyname, &category.Emergencyphone, &category.Image)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, category)

	}

	return ctgs, nil
}

// DeleteCategory removes a category from a database by its id
func (cri *CategoryRepositoryImpl) DeleteCategory(id int) error {

	_, err := cri.conn.Exec("DELETE FROM categories WHERE name = 'henock yonas'")
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}
