package test

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yahyrparedes/salva-template/cmd/test"
	"github.com/yahyrparedes/salva-template/pkg/models"
	"gorm.io/gorm"
	"testing"
)

func TestSetup(t *testing.T) {

	fmt.Println("TEST SETUP")
	app, db := test.InitializeConfigTest()
	SetupBrand(app, db)

}

func SetupBrand(app *fiber.App, db *gorm.DB) {
	db.AutoMigrate(&models.Brand{})
	brand := models.Brand{Name: "Go", Active: true}
	db.Create(&brand)
}
