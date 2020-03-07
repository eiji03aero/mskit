package log

import (
	"encoding/json"
	"log"

	"github.com/eiji03aero/mskit/utils"
)

func PrintJsonln(obj interface{}) {
	log.Println(SprintJson(obj))
}

func SprintJson(obj interface{}) string {
	o, _ := json.Marshal(obj)
	return string(o)
}

func PrintGet(obj interface{}) {
	_, name := utils.GetTypeName(obj)
	log.Println("get ", name, ": ", SprintJson(obj))
}

func PrintCreated(obj interface{}) {
	_, name := utils.GetTypeName(obj)
	log.Println(name, " created: ", SprintJson(obj))
}
