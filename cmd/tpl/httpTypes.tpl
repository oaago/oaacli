package {{.Package}}_{{.Method}}

type {{.UpPackage}}{{.UpMethod}}Service struct {
    {{.UpPackage}}{{.UpMethod}}Req
    {{.UpPackage}}{{.UpMethod}}Res
    //Dao      dao_oaauser.{{.UpPackage}}{{.UpMethod}}DaoType
    //DaoModel dao_oaauser.{{.UpPackage}}{{.UpMethod}}ModelType
    //DB       dao_oaauser.{{.UpPackage}}{{.UpMethod}}DaoType
}
//type {{.UpPackage}}{{.UpMethod}} dao_oaauser.{{.UpPackage}}{{.UpMethod}}ModelType
type {{.UpPackage}}{{.UpMethod}}Req struct{
    //{{.UpPackage}}{{.UpMethod}}
}

type {{.UpPackage}}{{.UpMethod}}Res struct{

}

func NewService{{.UpPackage}}() *{{.UpPackage}}{{.UpMethod}}Service {
	return &{{.UpPackage}}{{.UpMethod}}Service{
        //Dao:      dao_oaauser.{{.UpPackage}}{{.UpMethod}}Dao,
        //DaoModel: dao_oaauser.{{.UpPackage}}{{.UpMethod}}Model,
	}
}