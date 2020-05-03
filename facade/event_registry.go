package facade

import (
	"github.com/eiji03aero/mskit"
)

var (
	EventRegistry = mskit.NewEventRegistry()
)

func RegisterEvent(event interface{}) {
	EventRegistry.Set(event)
}
