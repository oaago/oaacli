package {{.Package}}
import (
	"context"
	rpc_{{.RpcName}}_{{.UpMethod}} "{{.Module}}/internal/api/rpc/{{.RpcName}}/{{.Method}}"
	"github.com/oaago/cloud/logx"
)
func (u *{{.UpPackage}}{{.UpMethod}}Type) {{.UpPackage}}{{.UpMethod}}Service(ctx context.Context, request *rpc_{{.RpcName}}_{{.UpMethod}}.{{.UpRpcName}}{{.UpMethod}}Request) (*rpc_{{.RpcName}}_{{.UpMethod}}.{{.UpRpcName}}{{.UpMethod}}Reply, error) {
	//TODO implement me
	logx.Logger.Info(request)
	// 调用其他rpc 服务的示例
	// 	res := rpc_fff_Ddd.NewRpcFffDddClient(rpc_fff_Ddd.RpcClientType{
	//		EtcdAddr:          "http://127.0.0.1:2379",
	//		RemoteServiceName: op.ConfigData.Server.Name,
	//	})
	//  c, cancel := context.WithTimeout(context.Background(), time.Second*2)
	//	resp, err := res.RpcFffDddService(c, &rpc_fff_Ddd.FffDddRequest{
	//		Name: "22222",
	//	})
	//  defer cancel()
	//	logx.Logger.Info(resp, err)
	return &rpc_{{.RpcName}}_{{.UpMethod}}.{{.UpRpcName}}{{.UpMethod}}Reply{}, nil
}