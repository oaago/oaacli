package {{.Package}}_{{.Method}}
{{if .HasTable}}
import (
    "{{.Module}}/internal/dao/dao_{{.DBName}}"
)
{{end}}

type {{.UpPackage}}{{.UpMethod}}Service struct {
    {{range $k,$v := .Met}}
    {{$v}}Req {{end}}
    {{range $k,$v := .Met}}
    {{$v}}Res {{end}}
    {{- if .HasTable}}Dao      dao_oaauser.{{.UpPackage}}{{.UpMethod}}DaoType{{end}}
}

func NewService{{.UpPackage}}{{.UpMethod}}() *{{.UpPackage}}{{.UpMethod}}Service {
	return &{{.UpPackage}}{{.UpMethod}}Service{
        {{if .HasTable}}Dao:  dao_oaauser.{{.UpPackage}}{{.UpMethod}}Dao, //nolint:typecheck{{end}}
	}
}