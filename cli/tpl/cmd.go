package tpl

func CmdAppTemplate() string {
	return `package main

import (
	"log"

	"{{ .PkgName }}/service"

	"github.com/eiji03aero/mskit/utils/logger"
)

func main() {
	_ = service.New()

	logger.Println(logger.CyanString("server started listening ..."))
}`
}

func CmdEnvTemplate() string {
	return `package main

import (
	"github.com/kelseyhightower/envconfig"
)

func init() {
	var env Env
	envconfig.Process("", &env)
}

type Env struct {
	PORT string
}`
}
