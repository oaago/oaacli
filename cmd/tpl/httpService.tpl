package {{.Package}}_{{.Method}}
import (
	"github.com/oaago/cloud/logx"
)
{{if eq .Upmet "Get"}}
func (u *{{.UpPackage}}{{.UpMethod}}Service) {{.Upmet}}{{.UpPackage}}{{.UpMethod}}Service() error {
	logx.Logger.Info("{{.Upmet}}{{.UpPackage}}{{.UpMethod}}Service")
	return nil
}
{{else if eq .Upmet "Delete"}}
func (u *{{.UpPackage}}{{.UpMethod}}Service) {{.Upmet}}{{.UpPackage}}{{.UpMethod}}Service() error {
	logx.Logger.Info("{{.Upmet}}{{.UpPackage}}{{.UpMethod}}Service")
	return nil
}
{{else}}
func (u *{{.UpPackage}}{{.UpMethod}}Service) {{.Upmet}}{{.UpPackage}}{{.UpMethod}}Service() ({{.UpPackage}}{{.UpMethod}}Res, error) {
	logx.Logger.Info("{{.Upmet}}{{.UpPackage}}{{.UpMethod}}Service")
	return u.{{.UpPackage}}{{.UpMethod}}Res,nil
}
{{end}}