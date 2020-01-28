package entity

type Category struct {
	ID             string `json:"id"`
	Name           string `json:"fullname"`
	Mothername     string `json:"mothername"`
	Image          string `json:"image"`
	AGE            string `json:"age"`
	Occupation     string `json:"occupation"`
	Relegion       string `json:"relegion"`
	Nationality    string `json:"natitionality"`
	Phonenumber    string `json:"phonenumber"`
	Sex            string `json:"sex"`
	Emergencyname  string `json:"emergencyname"`
	Emergencyphone string `json:"emergencyphone"`
}
type Childs struct {
	ChildName       string `json:"fullname"`
	ChildMothername string `json:"mothername"`
	CjildSex        string `json:"sex"`
	Birthdate       string `json:"emergencyname"`
}

// Role repesents application user roles
type Role struct {
	ID    uint
	Name  string `gorm:"type:varchar(255)"`
	Users []User
}

// User represents application user
type User struct {
	ID       uint
	FullName string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null; unique"`
	Phone    string `gorm:"type:varchar(100);not null; unique"`
	Password string `gorm:"type:varchar(255)"`
	RoleID   uint
	//Orders   []Order
}
type Session struct {
	ID         uint
	UUID       string `gorm:"type:varchar(255);not null"`
	Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}
