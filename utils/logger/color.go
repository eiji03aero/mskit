package logger

import (
	"github.com/fatih/color"
)

func RedString(s string, rest ...interface{}) string {
	return color.RedString(s, rest...)
}

func HiRedString(s string, rest ...interface{}) string {
	return color.HiRedString(s, rest...)
}

func BlueString(s string, rest ...interface{}) string {
	return color.BlueString(s, rest...)
}

func HiBlueString(s string, rest ...interface{}) string {
	return color.HiBlueString(s, rest...)
}

func CyanString(s string, rest ...interface{}) string {
	return color.CyanString(s, rest...)
}

func GreenString(s string, rest ...interface{}) string {
	return color.GreenString(s, rest...)
}

func YellowString(s string, rest ...interface{}) string {
	return color.YellowString(s, rest...)
}
