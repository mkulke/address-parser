//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=types.cfg.yaml ../spec.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=server.cfg.yaml ../spec.yaml

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	parser "github.com/openvenues/gopostal/parser"
)

func BuildResponse(components []parser.ParsedComponent) *Address {
	houseNumber := ""
	street := ""
	city := ""
	zipCode := ""
	state := ""
	for _, value := range components {
		if value.Label == "house_number" {
			houseNumber = value.Value
		}
		if value.Label == "road" {
			street = value.Value
		}
		if value.Label == "city" {
			city = value.Value
		}
		if value.Label == "postcode" {
			zipCode = value.Value
		}
		if value.Label == "state" {
			state = value.Value
		}
	}

	if houseNumber == "" || street == "" {
		return nil
	}

	address := Address{Street: street, HouseNumber: houseNumber}
	if city != "" {
		address.City = &city
	}
	if zipCode != "" {
		address.ZipCode = &zipCode
	}
	if state != "" {
		address.State = &state
	}

	return &address
}

type Api struct{}

func (a *Api) Parse(ctx echo.Context, params ParseParams) error {
	parsed := parser.ParseAddress(params.Q)
	res := BuildResponse(parsed)
	return ctx.JSON(http.StatusOK, res)
}
