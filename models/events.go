package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func AddEvent(request *CreateEvent) (*Event, error) {

	event := &Event{
		Title:       request.Title,
		Description: request.Description,
		Start:       request.Start,
		End:         request.End,
		Location:    request.Location,
		Latitude:    request.Latitude,
		Longitude:   request.Longitude,
	}
	ok, resp := event.Validate()

	if !ok {
		return nil, errors.New(resp)
	}
	GetDB().Create(event)
	if event.ID <= 0 {
		return nil, errors.New("failed to create event connection error")
	}

	return event, nil
}

func (eventToCheck *Event) Validate() (bool, string) { //not finished

	eventToCheck.Title = strings.Join(strings.Fields(eventToCheck.Title), " ")

	if len(eventToCheck.Title) < 2 || len(eventToCheck.Title) > 40 {
		return false, "Title must be 2-40 chars long"
	}

	eventToCheck.Description = strings.Join(strings.Fields(eventToCheck.Description), " ")

	if len(eventToCheck.Description) > 50 {
		return false, "The descrition must be less than 50 chars! "
	}

	eventToCheck.Location = strings.Join(strings.Fields(eventToCheck.Location), " ")

	if len(eventToCheck.Location) < 6 || len(eventToCheck.Location) > 40 {
		return false, "The Location field must be between 6-40 chars! "
	}

	lat, err := strconv.ParseFloat(eventToCheck.Latitude, 64)
	if err != nil {
		return false, "Latitude should be numeric value"
	}

	lng, err := strconv.ParseFloat(eventToCheck.Longitude, 64)
	if err != nil {
		return false, "Longitude should be numeric value"
	}

	ok, message := ValidateGeoCoords(lat, lng)
	if !ok {
		return false, message
	}

	const layout = "2006-02-01 15:04"
	start, err := time.Parse(layout, eventToCheck.Start)
	if err != nil {
		return false, "Wrong start data syntax"
	}

	end, err := time.Parse(layout, eventToCheck.End)
	if err != nil {
		return false, "Wrong end data syntax"
	}

	eventToCheck.Start = start.Format(layout)
	eventToCheck.End = end.Format(layout)

	if !start.Before(end) {
		return false, "Start of event must be before end!"
	}

	return true, "Requirement passed"
}

func GetEvent(eventId uint) (*Event, error) {

	event := &Event{}
	err := GetDB().Where("id = ?", eventId).First(event).Error
	if err != nil {
		return nil, err
	}
	//fmt.Printf("%T", err)

	return event, nil
}

func (eventToUpdate *Event) UpdateEventFields(updateFields *UpdateEvent) {
	//transport new values to event fields from update event structure
	if updateFields.Title != "" {
		eventToUpdate.Title = updateFields.Title
	}
	if updateFields.Description != "" {
		eventToUpdate.Description = updateFields.Description
	}
	if updateFields.Location != "" {
		eventToUpdate.Location = updateFields.Location
	}
	if updateFields.Latitude != "" {
		eventToUpdate.Latitude = updateFields.Latitude
	}
	if updateFields.Longitude != "" {
		eventToUpdate.Longitude = updateFields.Longitude
	}
	if updateFields.Start != "" {
		eventToUpdate.Start = updateFields.Start
	}
	if updateFields.End != "" {
		eventToUpdate.End = updateFields.End
	}

}

func UpdateEventRecord(updatedEventObject *Event) (*Event, error) {
	ok, resp := updatedEventObject.Validate()
	if !ok {
		return nil, errors.New(resp)
	}
	GetDB().Updates(updatedEventObject) //error handling

	return updatedEventObject, nil
}

func DeleteEvent(eventId uint) error {

	err := GetDB().Delete(&Event{}, eventId).Error
	if err != nil {
		return err
	}
	return nil
}

func FindAllEvents() (*[]Event, error) {
	//event := &Event{}
	var allEvents []Event
	result := GetDB().Find(&allEvents)

	fmt.Println(result.RowsAffected)

	if result.Error != nil {
		return nil, result.Error
	}
	return &allEvents, nil
}

func FindEventsInArea(latitude1, longitude1, latitude2, longitude2 string) (*[]Event, error) {

	lat1, err := strconv.ParseFloat(latitude1, 64)
	if err != nil {
		return nil, errors.New("latitude should be numeric value")
	}

	lat2, err := strconv.ParseFloat(latitude2, 64)
	if err != nil {
		return nil, errors.New("latitude should be numeric value")
	}

	lng1, err := strconv.ParseFloat(longitude1, 64)
	if err != nil {
		return nil, errors.New("longitude should be numeric value")
	}

	lng2, err := strconv.ParseFloat(longitude2, 64)
	if err != nil {
		return nil, errors.New("longitude should be numeric value")
	}

	ok, message := ValidateGeoCoords(lat1, lng1)
	if !ok {
		return nil, errors.New(message)
	}
	ok, message = ValidateGeoCoords(lat2, lng2)
	if !ok {
		return nil, errors.New(message)
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

	var allEvents []Event
	var allEventsInArea []Event
	result := GetDB().Find(&allEvents)

	if result.Error != nil {
		return nil, result.Error
	}
	for _, elem := range allEvents {

		lat, _ := strconv.ParseFloat(elem.Latitude, 64)
		lng, _ := strconv.ParseFloat(elem.Longitude, 64)

		if lat <= top && lat >= bottom && lng >= left && lng <= right {
			allEventsInArea = append(allEventsInArea, elem)
		}
	}

	return &allEventsInArea, nil
}

func ValidateGeoCoords(lat, lng float64) (bool, string) {

	if lat < -90 || lat > 90 {
		return false, "Latitude is out of bounds"
	}
	if lng < -180 || lng > 180 {
		return false, "Longitude is out of bounds"
	}
	return true, "coords validated"
}
