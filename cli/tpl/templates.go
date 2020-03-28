package tpl

func Interface() string {
	return `
type {{ .InterfaceName }} interface {
}`
}
