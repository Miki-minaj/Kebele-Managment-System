package menu

import "github.com/miki-minaj/Kebele-Managment-System/entity"

// Service service
type Service interface {
	StoreInformations(informations entity.Infos)
}
