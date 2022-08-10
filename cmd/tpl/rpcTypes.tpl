package {{.Package}}
import (
	{{.UpPackage}}{{.UpMethod}} "{{.Module}}/internal/api/rpc/{{.Package}}/{{.Method}}"
)
type {{.UpPackage}}{{.UpMethod}}Type struct {
	{{.UpPackage}}{{.UpMethod}}.Unimplemented{{.UpPackage}}{{.UpMethod}}Server
}

func NewService{{.UpPackage}}{{.UpMethod}}() *{{.UpPackage}}{{.UpMethod}}Type {
	return &{{.UpPackage}}{{.UpMethod}}Type{}
}