package mskit

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/eiji03aero/mskit/utils"
)

// EventMap is a map to hold Event structs
type EventMap map[string]reflect.Type

// EventRegistry provides feature to manage EventMap
type EventRegistry struct {
	events EventMap
}

// NewEventRegistry creates new struct
func NewEventRegistry() *EventRegistry {
	return &EventRegistry{
		events: EventMap{},
	}
}

// Set registers new struct
func (er *EventRegistry) Set(event interface{}) error {
	rawType, name := utils.GetType(event)
	er.events[name] = rawType
	return nil
}

// Get returns pointer to the struct that was registered by Set
func (er *EventRegistry) Get(name string) (interface{}, error) {
	rawType, ok := er.events[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("not registered event: %s", name))
	}
	return reflect.New(rawType).Interface(), nil
}
