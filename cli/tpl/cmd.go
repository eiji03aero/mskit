package tpl

func CmdAppTemplate() string {
	return `package main

import (
	"log"
	"{{ .PkgName }}/service"
)

func main() {
	_ = service.New()

	log.Println("server started listening ...")
}`
}

func CmdEnvTemplate() string {
	return `package main

import (
	"github.com/keyseyhightower/envconfig"
)

func init() {
	var env Env
	envconfig.Process("", &env)
}

type Env struct {
	PORT string
}`
}
