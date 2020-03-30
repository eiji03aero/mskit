package tpl

func CmdAppTemplate() string {
	return `package main

import (
	"{{ .PkgName }}/service"
)

func main() {
	_ := service.New()

	bff := make(chan bool)
	<-bff
}`
}
