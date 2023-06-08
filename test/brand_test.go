package test

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/stretchr/testify/assert"
	"github.com/yahyrparedes/salva-template/cmd/test"
	"github.com/yahyrparedes/salva-template/pkg/mappers"
	"io"
	"net/http/httptest"
	"testing"
)

func TestBrands(t *testing.T) {
	app, _ := test.InitializeConfigTest()

	req := httptest.NewRequest(fiber.MethodGet, "/api/v1/brands", nil)
	res, err := app.Test(req, 10)

	// Test status
	utils.AssertEqual(t, nil, err, "Ups Error!")
	utils.AssertEqual(t, 200, res.StatusCode, "Status Code")

	body, _ := io.ReadAll(res.Body)
	jsonBody := string(body)
	var brands []mappers.Brand
	json.Unmarshal([]byte(jsonBody), &brands)
	var id uint = 1

	// Test response
	assert.Equalf(t, 1, len(brands), "Brands Size!")
	assert.Equalf(t, id, brands[0].Id, "Brands ID!")
	assert.Equalf(t, "Go", brands[0].Name, "Brands Name!")
}

func TestBrandDetail(t *testing.T) {
	// t.Parallel()
	app, _ := test.InitializeConfigTest()
	req := httptest.NewRequest(fiber.MethodGet, "/api/v1/brand/1", nil)
	res, err := app.Test(req, -1)

	// Test status
	utils.AssertEqual(t, nil, err, "Ups Error!")
	utils.AssertEqual(t, 200, res.StatusCode, "Status Code")

	body, _ := io.ReadAll(res.Body)
	jsonBody := string(body)
	var brand mappers.Brand
	json.Unmarshal([]byte(jsonBody), &brand)
	var id uint = 1

	// Test response
	assert.Equalf(t, id, brand.Id, "Brand ID!")
	assert.Equalf(t, "Go", brand.Name, "Brand Name!")
	assert.Equalf(t, true, brand.Active, "Brand Active!")
}
