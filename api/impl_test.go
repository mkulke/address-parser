package api

import (
	"testing"

	parser "github.com/openvenues/gopostal/parser"
)

func TestEmpty(t *testing.T) {
	var parsed []parser.ParsedComponent
	res := BuildResponse(parsed)
	if res != nil {
		t.Errorf("expected nil, got address")
	}
}

func TestNoHouseNumber(t *testing.T) {
	parsed := []parser.ParsedComponent{
		parser.ParsedComponent{Label: "road", Value: "Some Road"},
	}
	res := BuildResponse(parsed)
	if res != nil {
		t.Errorf("got address, expected nil")
	}
}

func TestNoRoad(t *testing.T) {
	parsed := []parser.ParsedComponent{
		parser.ParsedComponent{Label: "house_number", Value: "123"},
	}
	res := BuildResponse(parsed)
	if res != nil {
		t.Errorf("got address, expected nil")
	}
}

func TestSimple(t *testing.T) {
	road := "Some Road"
	no := "123"
	parsed := []parser.ParsedComponent{
		parser.ParsedComponent{Label: "road", Value: road},
		parser.ParsedComponent{Label: "house_number", Value: no},
	}
	res := BuildResponse(parsed)
	if res == nil {
		t.Errorf("got nil, expected address")
	}
	if res.Street != road {
		t.Errorf("expected %s, got %s", road, res.Street)
	}
	if res.HouseNumber != no {
		t.Errorf("expected %s, got %s", no, res.HouseNumber)
	}
}

func TestFull(t *testing.T) {
	city := "Paris"
	zipCode := "12345"
	state := "California"

	parsed := []parser.ParsedComponent{
		parser.ParsedComponent{Label: "road", Value: "some road"},
		parser.ParsedComponent{Label: "house_number", Value: "123"},
		parser.ParsedComponent{Label: "city", Value: city},
		parser.ParsedComponent{Label: "postcode", Value: zipCode},
		parser.ParsedComponent{Label: "state", Value: state},
	}
	res := BuildResponse(parsed)
	if res == nil {
		t.Errorf("got nil, expected address")
	}
	if *res.City != city {
		t.Errorf("expected %s, got %s", city, *res.City)
	}
	if *res.ZipCode != zipCode {
		t.Errorf("expected %s, got %s", zipCode, *res.ZipCode)
	}
	if *res.State != state {
		t.Errorf("expected %s, got %s", state, *res.State)
	}
}
