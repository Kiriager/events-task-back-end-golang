package models

import (
	"errors"
	"fmt"
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

	newEvent := &Event{
		Title:       eventData.Title,
		Description: eventData.Description,
		Start:       eventData.Start,
		End:         eventData.End,
		LocationID:  eventData.LocationID,
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

	newEventLocation, err := GetLocation(eventToCheck.LocationID)
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

	err := GetDB().Where("id = ?", eventId).First(event).Error
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

	if updateData.Start.IsZero() != true {
		eventToUpdate.Start = updateData.Start
	}
	if updateData.End.IsZero() != true {
		eventToUpdate.End = updateData.End
	}
	eventToUpdate.LocationID = updateData.LocationID

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

	err = GetDB().Updates(eventToUpdate).Where("id = ?", eventId).Error
	if err != nil {
		return nil, err
	}

	updatedEventRecord, err := GetEvent(*eventId)
	if err != nil {
		return nil, err
	}

	return updatedEventRecord, nil
}

/*
func (updateEventData *UpdateEvent) validateUpdateEventData() (*Event, error) {
	updateEventData.Title = strings.Join(strings.Fields(updateEventData.Title), " ")

	if updateEventData.Title != "" {
		if len(updateEventData.Title) < 4 || len(updateEventData.Title) > 40 {
			return nil, errors.New("title must be 4-40 chars long")
		}
	}

	updateEventData.Description = strings.Join(strings.Fields(updateEventData.Description), " ")

	if updateEventData.Description != "" {
		if len(updateEventData.Description) > 50 {
			return nil, errors.New("descrition must be less than 50 chars")
		}
	}

	if updateEventData.LocationID != 0 {
		newEventLocation, err := GetLocation(updateEventData.LocationID)
		if err != nil {
			return nil, err
		}

		if newEventLocation.ID == 0 {
			return nil, errors.New("location not found")
		}
	}

	if !eventDataToCheck.End.Before(eventDataToCheck.End) {
		return nil, errors.New("start of event must be before end")
	}

	newEvent := &Event{
		Title:       eventDataToCheck.Title,
		Description: eventDataToCheck.Description,
		Start:       eventDataToCheck.Start,
		End:         eventDataToCheck.End,
		LocationID:  eventDataToCheck.LocationID,
	}

	return newEvent, nil
}
*/

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

	fmt.Println(result.RowsAffected)

	if result.Error != nil {
		return nil, result.Error
	}
	return &allEvents, nil
}

/*
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
*/
func ValidateGeoCoords(lat, lng float64) error {

	if lat < -90 || lat > 90 {
		return errors.New("latitude is out of bounds")
	}
	if lng < -180 || lng > 180 {
		return errors.New("longitude is out of bounds")
	}
	return nil
}
