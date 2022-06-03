package models

import (
	"errors"
	"strconv"
	"strings"
)

func RecordLocation(locationData *RegisterLocation) (*Location, error) {
	err, newLocation := locationData.ValidateNewLocation()

	if err != nil {
		return nil, err
	}

	GetDB().Create(newLocation)

	if newLocation.ID <= 0 {
		return nil, errors.New("failed to create event connection error")
	}

	return newLocation, nil
}

func (locationData *RegisterLocation) ValidateNewLocation() (error, *Location) { //not finished

	locationData.Title = strings.Join(strings.Fields(locationData.Title), " ")

	if len(locationData.Title) < 2 || len(locationData.Title) > 40 {
		return errors.New("Title of location must be 4-20 chars! "), nil
	}

	locationData.Description = strings.Join(strings.Fields(locationData.Description), " ")

	if len(locationData.Description) > 50 {
		return errors.New("The descrition must be less than 50 chars! "), nil
	}

	lat, err := strconv.ParseFloat(locationData.Latitude, 64)
	if err != nil {
		return err, nil
	}

	lng, err := strconv.ParseFloat(locationData.Longitude, 64)
	if err != nil {
		return err, nil
	}

	ok, message := ValidateGeoCoords(lat, lng)
	if !ok {
		return errors.New(message), nil
	}

	validatedLocation := &Location{
		Title:       locationData.Title,
		Description: locationData.Description,
		Latitude:    lat,
		Longitude:   lng,
	}

	return nil, validatedLocation
}
