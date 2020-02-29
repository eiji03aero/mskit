package utils

import (
	"fmt"
)

func NewErrNotEnoughPropertiesSet(args [][]interface{}) error {
	var propertiesString string
	for _, arg := range args {
		pairString := fmt.Sprintf(" %s=%v", arg[0], arg[1])
		propertiesString += pairString
	}

	return fmt.Errorf("Not enough properties set: %s", propertiesString)
}
