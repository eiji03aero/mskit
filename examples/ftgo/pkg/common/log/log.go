package log

import (
	"encoding/json"
	"log"
)

func PrintJsonln(obj interface{}) {
	log.Println(SprintJson(obj))
}

func SprintJson(obj interface{}) string {
	o, _ := json.Marshal(obj)
	return string(o)
}
