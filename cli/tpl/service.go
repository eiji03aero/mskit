package tpl

func RootService() string {
	return `package {{ .PkgName }}

type Service interface {
}`
}

func ServiceTemplate() string {
	return `package service

import (
	"{{ .PkgName }}"
)

type service struct {
}

func New() {{ .PkgName }}.Service {
	return &service{}
}`
}
