package models

import (
	"errors"
	"strconv"
	"strings"
)

func RecordNewEvent(newEventData *RegisterEvent) (*Event, error) {
	newEvent := newEventData.constructEvent()

	err := newEvent.ValidateEvent()
	if err != nil {
		return nil, err
	}

	err = GetDB().Create(newEvent).Error
	if err != nil {
		return nil, err
	}

	if newEvent.ID <= 0 {
		return nil, errors.New("failed to create event connection error")
	}

	return newEvent, nil
}

func (eventData *RegisterEvent) constructEvent() *Event {
	//location, _ := GetLocation(eventData.LocationID)

	newEvent := &Event{
		Title:       eventData.Title,
		Description: eventData.Description,
		Start:       eventData.Start,
		End:         eventData.End,
		LocationId:  eventData.LocationId,
		//Location:    *location,
	}

	return newEvent
}

func (eventToCheck *Event) ValidateEvent() error {
	eventToCheck.Title = strings.Join(strings.Fields(eventToCheck.Title), " ")
	if len(eventToCheck.Title) < 4 || len(eventToCheck.Title) > 40 {
		return errors.New("title must be 4-40 chars long")
	}

	eventToCheck.Description = strings.Join(strings.Fields(eventToCheck.Description), " ")
	if len(eventToCheck.Description) > 50 {
		return errors.New("descrition must be less than 50 chars")
	}

	newEventLocation, err := GetLocation(eventToCheck.LocationId)
	if err != nil {
		return err
	}

	if newEventLocation.ID == 0 {
		return errors.New("location not found")
	}

	if !eventToCheck.Start.Before(eventToCheck.End) {
		return errors.New("start of event must be before end")
	}

	return nil
}

func GetEvent(eventId uint) (*Event, error) {
	event := &Event{}

	err := GetDB().Where("id = ?", eventId).Preload("Location").First(event).Error
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (eventToUpdate *Event) UpdateEventFields(updateData *UpdateEvent) {

	if updateData.Title != "" {
		eventToUpdate.Title = updateData.Title
	}
	if updateData.Description != "" {
		eventToUpdate.Description = updateData.Description
	}

	if !updateData.Start.IsZero() {
		eventToUpdate.Start = updateData.Start
	}
	if !updateData.End.IsZero() {
		eventToUpdate.End = updateData.End
	}
	if updateData.LocationId != 0 {
		eventToUpdate.LocationId = updateData.LocationId
	}

}

func UpdateEventRecord(updateEventData *UpdateEvent, eventId *uint) (*Event, error) {

	eventToUpdate, err := GetEvent(*eventId)
	if err != nil {
		return nil, err
	}

	eventToUpdate.UpdateEventFields(updateEventData)

	err = eventToUpdate.ValidateEvent()
	if err != nil {
		return nil, err
	}

	err = GetDB().Updates(eventToUpdate).Where("id = ?", *eventId).Error
	if err != nil {
		return nil, err
	}

	updatedEventRecord, err := GetEvent(*eventId)
	if err != nil {
		return nil, err
	}

	return updatedEventRecord, nil
}

func DeleteEvent(eventId uint) error {
	err := GetDB().Delete(&Event{}, eventId).Error
	if err != nil {
		return err
	}
	return nil
}

func FindAllEvents() (*[]Event, error) {
	var allEvents []Event
	result := GetDB().Find(&allEvents)

	//fmt.Println(result.RowsAffected)

	if result.Error != nil {
		return nil, result.Error
	}
	return &allEvents, nil
}

func FindAllEventsInLocation(locationId uint) ([]Event, error) {
	location, err := GetLocation(locationId)
	if err != nil {
		return nil, err
	}
	events := []Event{}
	// GetDB().Debug().Preload("Events").Find(&location)
	// events = location.Events
	GetDB().Model(&location).Debug().Association("Events").Find(&events)
	return events, nil
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

	var allEventsInArea []Event
	var allLocations []Location

	err = GetDB().Find(&allLocations).Error
	if err != nil {
		return nil, err
	}

	for _, elem := range allLocations {
		inArea, err := elem.IsInArea(lat1, lng1, lat2, lng2)
		if err != nil {
			return nil, err
		}

		if inArea {
			eventsInLocation := []Event{}
			GetDB().Model(&elem).Debug().Association("Events").Find(&eventsInLocation)
			for _, event := range eventsInLocation {
				GetDB().Debug().Preload("Location").Find(&event)
				allEventsInArea = append(allEventsInArea, event)
			}
		}
	}

	return &allEventsInArea, nil
}

func ValidateGeoCoords(lat, lng float64) error {

	if lat < -90 || lat > 90 {
		return errors.New("latitude is out of bounds")
	}
	if lng < -180 || lng > 180 {
		return errors.New("longitude is out of bounds")
	}
	return nil
}
