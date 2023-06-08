package services

import (
	"github.com/yahyrparedes/salva-template/pkg/mappers"
	QueryBrand "github.com/yahyrparedes/salva-template/pkg/repositories"
)

func GetBrands() []mappers.Brand {
	brands := QueryBrand.FindAll()
	mBrands := mappers.RequestBrands(brands)
	return mBrands
}

func GetBrand(id int) mappers.Brand {
	brand := QueryBrand.FindByID(id)
	mBrand := mappers.RequestBrand(brand)
	return mBrand
}
