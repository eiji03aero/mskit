package mskit

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/eiji03aero/mskit/utils"
)

type EventMap map[string]reflect.Type

type EventRegistry struct {
	events EventMap
}

func NewEventRegistry() *EventRegistry {
	return &EventRegistry{
		events: EventMap{},
	}
}

func (er *EventRegistry) Set(event interface{}) error {
	rawType, name := utils.GetTypeName(event)
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
