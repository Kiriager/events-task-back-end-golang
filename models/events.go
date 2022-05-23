package models

import (
	"errors"
)

func AddEvent(request *CreateEvent) (*Event, error) {

	event := &Event{Title: request.Title}
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

func (account *Event) Validate() (bool, string) {

	if len(account.Title) < 6 {
		return false, "Password is required"
	}

	return true, "Requirement passed"
}
