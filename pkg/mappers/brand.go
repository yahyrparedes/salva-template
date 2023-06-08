package mappers

import (
	"fmt"
	"github.com/yahyrparedes/salva-template/pkg/models"
)

type Brand struct {
	Id     uint
	Name   string
	Active bool
}

func RequestBrands(brands []models.Brand) []Brand {
	fmt.Println(brands)
	var bb []Brand
	for i, b := range brands {
		fmt.Println(i, b)
		bb = append(bb, RequestBrand(b))
	}

	return bb
}

func RequestBrand(brand models.Brand) Brand {
	var b Brand
	b.Id = brand.ID
	b.Name = brand.Name
	b.Active = brand.Active
	return b
}
