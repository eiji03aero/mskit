package logger

import (
	"encoding/json"
)

func formatArgs(args []interface{}) (result []interface{}) {
	for _, arg := range args {
		a := formatDataToJson(arg)
		result = append(result, a)
	}
	return
}

func formatArgsRest(args []interface{}, rest []interface{}) (result []interface{}) {
	args = append(args, rest...)
	return formatArgs(args)
}

func formatDataToJson(data interface{}) (result interface{}) {
	switch a := data.(type) {
	case string:
		result = a
	case []byte:
		result = string(a)
	default:
		aJson, err := json.Marshal(a)
		if err != nil {
			panic(err)
		}
		result = string(aJson)
	}

	return
}
