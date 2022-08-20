package {{.Package}}
import (
	"github.com/oaago/cloud/logx"
)

func (Srv *{{.UpPackage}}Service) {{.UpMethod}}Service() error {
	logx.Logger.Info("{{.UpMethod}}Service")
	// 针对业务直接操作 Srv.{{.UpPackage}}{{.UpMethod}}Res 返回结构即可
	return nil
}