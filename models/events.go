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

	GetDB().Create(event)

	if event.ID <= 0 {
		return nil, errors.New("failed to create event connection error")
	}

	return event, nil
}

func (eventToCheck *Event) Validate() (bool, string) {
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
<<<<<<< HEAD
		fmt.Println("Hello")
=======
>>>>>>> 9ee379b9b7b4d7d0a3f79b3e8ae4ee858d0bb8a2
	}

	return isOk, "Requirement passed"
}
