package menu

import "github.com/miki-minaj/Kebele-Managment-System/entity"

// Repository specifies informaions list
type Repository interface {
	StoreInformations(informations entity.Infos)
}
