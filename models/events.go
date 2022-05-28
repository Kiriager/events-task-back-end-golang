package models

import (
	"errors"
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

	if lat < -90 || lat > 90 {
		return false, "Latitude is out of bounds"
	}

	lng, err := strconv.ParseFloat(eventToCheck.Longitude, 64)
	if err != nil {
		return false, "Longitude should be numeric value"
	}

	if lng < -180 || lng > 180 {
		return false, "Longitude is out of bounds"
	}

	start, err := time.Parse("2006-01-02 15:04", eventToCheck.Start)
	if err != nil {
		return false, "Wrong start data syntax"
	}

	end, err := time.Parse("2006-01-02 15:04", eventToCheck.End)
	if err != nil {
		return false, "Wrong end data syntax"
	}

	eventToCheck.Start = start.Format("2022-02-24 04:15")
	eventToCheck.End = end.Format("2022-02-24 04:15")

	if !start.Before(end) {
		return false, "Start of event must be before end!"
	}

	return true, "Requirement passed"
}

func GetEvent(eventId uint) *Event {

	event := &Event{}
	GetDB().Where("id = ?", eventId).First(event)

	return event
}

func (eventToUpdate *Event) UpdateEventFields(updateFields *UpdateEvent) {
	//transport new values to event fields from update event structure
	eventToUpdate.Title = updateFields.Title
	//ok, resp := eventToUpdate.Validate()
	/*
		if !ok {
			return nil, errors.New(resp)
		}*/
}

func (eventToUpdate *Event) UpdateEventRecord() *Event {
	GetDB().Updates(eventToUpdate)
	return eventToUpdate
}

func DeleteEvent(eventId uint) {

	//GetDB().Delete(eventToDelete)
	GetDB().Delete(&Event{}, eventId)
}
