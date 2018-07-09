// Package calendar contains operations which are performed on a general calendar
package calendar

import (
	"fmt"
	"sort"
)

// OverlappingEvent represents a pair of events that overlap with each other
type OverlappingEvent struct {
	// Event1 specifies the first event that overlaps with another event
	Event1 Event

	// Event2 specifies the second event that overlaps with first event
	Event2 Event
}

// GetOverlappingEvents returns all pairs of overlapping events
func GetOverlappingEvents(events []Event) ([]OverlappingEvent, error) {
	var overlappingEvents []OverlappingEvent
	if len(events) < 2 {
		return overlappingEvents, fmt.Errorf("events param must contain at least 2 events")
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].StartTime < events[j].StartTime
	})

	for i := 1; i < len(events); i++ {
		if areOverlapping(events[i-1], events[i]) {
			overlappingEvent := OverlappingEvent{
				Event1: events[i-1],
				Event2: events[i],
			}
			overlappingEvents = append(overlappingEvents, overlappingEvent)
		}
	}

	return overlappingEvents, nil
}

// areOverlapping determines whether two events overlap with each other or not
func areOverlapping(event1, event2 Event) bool {
	return event2.StartTime < event1.EndTime
}
