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

func (locationData *RegisterLocation) ValidateNewLocation() (*Location, error) { //not finished

	locationData.Title = strings.Join(strings.Fields(locationData.Title), " ")

	if len(locationData.Title) < 2 || len(locationData.Title) > 40 {
		return nil, errors.New("Title of location must be 4-20 chars! ")
	}

	locationData.Description = strings.Join(strings.Fields(locationData.Description), " ")

	if len(locationData.Description) > 50 {
		return nil, errors.New("The descrition must be less than 50 chars! ")
	}
	/*
		lat, err := strconv.ParseFloat(locationData.Latitude, 64)
		if err != nil {
			return err, nil
		}

		lng, err := strconv.ParseFloat(locationData.Longitude, 64)
		if err != nil {
			return err, nil
		}
	*/
	lat := locationData.Latitude
	lng := locationData.Longitude

	ok, message := ValidateGeoCoords(lat, lng)
	if !ok {
		return nil, errors.New(message)
	}

	validatedLocation := &Location{
		Title:       locationData.Title,
		Description: locationData.Description,
		Latitude:    lat,
		Longitude:   lng,
	}

	return validatedLocation, nil
}
