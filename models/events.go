package models

import (
	"errors"
	"fmt"
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

	if event.ID <= 0 {
		return nil, errors.New("failed to create event connection error")
	}

	GetDB().Create(event)
	return event, nil
}

func (eventToCheck *Event) Validate() (bool, string) { //not finished
	errs := ""
	var isOk bool = true
	if eventToCheck.Title == "" {
		isOk = false
		errs += "The title is required! "
	}

	if len(eventToCheck.Title) < 2 || len(eventToCheck.Title) > 40 {
		isOk = false
		errs += "The title field must be between 2-40 chars! "
	}

	if isOk {
		errs = "Requirement passed"
		fmt.Println("Event is valid")
		fmt.Println("Hello")
	}

	return isOk, "Requirement passed"
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
