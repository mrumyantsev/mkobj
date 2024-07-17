package templates

const TemplateObject = `package %s

type %s struct {
}

func New() *%s {
	return &%s{}
}
`
