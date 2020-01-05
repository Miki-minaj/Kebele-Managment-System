package menu

import "github.com/miki-minaj/Kebele-Managment-System/entity"

//catagory repository
type CategoryRepository interface {
	StoreInformations(informations entity.Infos)
}
