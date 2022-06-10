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

func (location *Location) IsInArea(lat1, lng1, lat2, lng2 float64) (bool, error) {
	err := ValidateGeoCoords(lat1, lng1)
	if err != nil {
		return false, err
	}
	err = ValidateGeoCoords(lat2, lng2)
	if err != nil {
		return false, err
	}

	var top, bottom, left, right float64

	if lat1 > lat2 {
		top = lat1
		bottom = lat2
	} else {
		top = lat2
		bottom = lat1
	}
	if lng1 < lng2 {
		left = lng1
		right = lng2
	} else {
		left = lng2
		right = lng1
	}

	if location.Latitude <= top && location.Latitude >= bottom && location.Longitude >= left && location.Longitude <= right {
		return true, nil
	}
	return false, nil
}

func FindAllLocations() (*[]Location, error) {
	var allLocations []Location
	result := GetDB().Find(&allLocations)

	if result.Error != nil {
		return nil, result.Error
	}
	return &allLocations, nil
}

func UpdateLocationRecord(updateLocationData *UpdateLocation, locationId *uint) (*Location, error) {

	LocationToUpdate, err := GetLocation(*locationId)
	if err != nil {
		return nil, err
	}

	LocationToUpdate.UpdateLocationFields(updateLocationData)

	err = LocationToUpdate.ValidateLocation()
	if err != nil {
		return nil, err
	}

	err = GetDB().Updates(LocationToUpdate).Where("id = ?", *locationId).Error
	if err != nil {
		return nil, err
	}
	//bug: lat = 0, lng = 0 must be possible

	updatedLocationRecord, err := GetLocation(*locationId)
	if err != nil {
		return nil, err
	}

	return updatedLocationRecord, nil
}

func (locationToUpdate *Location) UpdateLocationFields(updateData *UpdateLocation) {

	if updateData.Title != "" {
		locationToUpdate.Title = updateData.Title
	}
	if updateData.Description != "" {
		locationToUpdate.Description = updateData.Description
	}

	locationToUpdate.Latitude = updateData.Latitude
	locationToUpdate.Longitude = updateData.Longitude

}

func (location *Location) ValidateLocation() error {

	location.Title = strings.Join(strings.Fields(location.Title), " ")

	if len(location.Title) < 2 || len(location.Title) > 40 {
		return errors.New("title of location must be 4-20 chars! ")
	}

	location.Description = strings.Join(strings.Fields(location.Description), " ")

	if len(location.Description) > 50 {
		return errors.New("descrition must be less than 50 chars! ")
	}

	err := ValidateGeoCoords(location.Latitude, location.Longitude)
	if err != nil {
		return err
	}

	return nil
}

func DeleteLocation(locationId uint) error {
	_, err := GetLocation(locationId)
	if err != nil {
		return err
	}

	err = GetDB().Delete(&Location{}, locationId).Error
	if err != nil {
		return err
	}
	return nil
}
