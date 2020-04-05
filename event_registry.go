package mskit

import (
	"fmt"
	"reflect"

	"github.com/eiji03aero/mskit/utils"
	"github.com/eiji03aero/mskit/utils/logger"
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
func (er *EventRegistry) Get(name string) (event interface{}, err error) {
	rawType, ok := er.events[name]
	if !ok {
		err = fmt.Errorf("Not registered event %s", name)
		logger.Println(logger.RedString(err.Error()))
		return
	}

	event = reflect.New(rawType).Interface()
	return
}
