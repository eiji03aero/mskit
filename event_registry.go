package mskit

import (
	"errors"
	"fmt"
	"github.com/eiji03aero/mskit/utils"
	"reflect"
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

func (er *EventRegistry) Get(name string) (interface{}, error) {
	rawType, ok := er.events[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("not registered event: %s", name))
	}
	return reflect.New(rawType).Interface(), nil
}
