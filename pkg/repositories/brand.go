package repositories

import (
	"github.com/yahyrparedes/salva-template/cmd/database"
	"github.com/yahyrparedes/salva-template/pkg/models"
)

func FindAll() []models.Brand {
	var brands []models.Brand
	//database.Connection.Table("brand").Find(&brands)
	database.Connection.Find(&brands)
	return brands
}

func FindByID(id int) models.Brand {
	var brand models.Brand
	database.Connection.Find(&brand, "id", id)
	return brand
}
