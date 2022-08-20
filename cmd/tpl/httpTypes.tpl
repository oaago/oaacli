package {{.Package}}_{{.Method}}

import (
    "{{.Module}}/internal/dao/dao_{{.DBName}}"
)

type {{.UpPackage}}{{.UpMethod}}Service struct {
    {{range $k,$v := .Met}}
    {{$v}}Req {{end}}
    {{range $k,$v := .Met}}
    {{$v}}Res {{end}}
    Dao      dao_oaauser.{{.UpPackage}}{{.UpMethod}}DaoType
}

func NewService{{.UpPackage}}{{.UpMethod}}() *{{.UpPackage}}{{.UpMethod}}Service {
	return &{{.UpPackage}}{{.UpMethod}}Service{
        Dao:  dao_oaauser.{{.UpPackage}}{{.UpMethod}}Dao,
	}
}