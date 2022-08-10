package {{.Package}}

type {{.UpPackage}}{{.UpMethod}}Service struct {
    {{.UpPackage}}{{.UpMethod}}Req
    {{.UpPackage}}{{.UpMethod}}Res
}

type {{.UpPackage}}{{.UpMethod}}Req struct{

}

type {{.UpPackage}}{{.UpMethod}}Res struct{

}

type {{.UpPackage}}{{.UpMethod}}IDAL interface {
    Put{{.UpPackage}}{{.UpMethod}}Service(req {{.UpPackage}}{{.UpMethod}}Req)
    Get{{.UpPackage}}{{.UpMethod}}Service(id int64)
    Update{{.UpPackage}}{{.UpMethod}}Service(req {{.UpPackage}}{{.UpMethod}}Req)
    Delete{{.UpPackage}}{{.UpMethod}}Service(ids []int64)
}

func NewService{{.UpPackage}}() *{{.UpPackage}}{{.UpMethod}}Service {
	return &{{.UpPackage}}{{.UpMethod}}Service{}
}