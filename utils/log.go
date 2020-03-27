package utils

import (
	"encoding/json"
	"log"
)

func PrintJsonln(obj interface{}) {
	log.Println(SprintJson(obj))
}

func PrintlnWithJson(str string, obj interface{}) {
	log.Println(str, SprintJson(obj))
}

func SprintJson(obj interface{}) string {
	o, _ := json.Marshal(obj)
	return string(o)
}

func PrintGet(obj interface{}) {
	_, name := GetTypeName(obj)
	log.Println("get ", name, ": ", SprintJson(obj))
}

func PrintCreated(obj interface{}) {
	_, name := GetTypeName(obj)
	log.Println(name, " created: ", SprintJson(obj))
}
