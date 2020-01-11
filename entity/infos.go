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
