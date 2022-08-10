package {{.Package}}_{{.Method}}

type {{.UpPackage}}{{.UpMethod}}Service struct {
    {{.UpPackage}}{{.UpMethod}}Req
    {{.UpPackage}}{{.UpMethod}}Res
}

type {{.UpPackage}}{{.UpMethod}}Req struct{

}

type {{.UpPackage}}{{.UpMethod}}Res struct{

}

func NewService{{.UpPackage}}() *{{.UpPackage}}{{.UpMethod}}Service {
	return &{{.UpPackage}}{{.UpMethod}}Service{}
}