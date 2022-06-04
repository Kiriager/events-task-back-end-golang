package models

import (
	"errors"
	"strings"
)

func RecordLocation(locationData *RegisterLocation) (*Location, error) {
	newLocation, err := locationData.ValidateNewLocation()

	if err != nil {
		return nil, err
	}

	GetDB().Create(newLocation)

	if newLocation.ID <= 0 {
		return nil, errors.New("failed to create event connection error")
	}

	return newLocation, nil
}

func (locationData *RegisterLocation) ValidateNewLocation() (*Location, error) {

	locationData.Title = strings.Join(strings.Fields(locationData.Title), " ")

	if len(locationData.Title) < 2 || len(locationData.Title) > 40 {
		return nil, errors.New("title of location must be 4-20 chars! ")
	}

	locationData.Description = strings.Join(strings.Fields(locationData.Description), " ")

	if len(locationData.Description) > 50 {
		return nil, errors.New("descrition must be less than 50 chars! ")
	}

	lat := locationData.Latitude
	lng := locationData.Longitude

	err := ValidateGeoCoords(lat, lng)
	if err != nil {
		return nil, err
	}

	validatedLocation := &Location{
		Title:       locationData.Title,
		Description: locationData.Description,
		Latitude:    lat,
		Longitude:   lng,
	}

	return validatedLocation, nil
}

func GetLocation(locationId uint) (*Location, error) {

	location := &Location{}
	err := GetDB().Where("id = ?", locationId).First(location).Error
	if err != nil {
		return nil, err
	}

	return location, nil
}
