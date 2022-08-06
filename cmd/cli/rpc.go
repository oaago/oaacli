package cli

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/oaago/oaago/cmd/tpl"
	"github.com/oaago/oaago/utils"
)

type RpcList struct {
	Module    string
	UpModule  string
	Package   string
	UpMethod  string
	UpPackage string
	Method    string
}

type RpcRouter struct {
	RPCMap    []RpcList
	Package   string
	UpPackage string
	Module    string
	UpModule  string
	UpMethod  string
	Method    string
}

var RpcRoute RpcRouter

func genProto(args []string, path string) {
	arg := strings.Split(args[0], "/")
	fmt.Println(arg)
	var protoFile *os.File
	var err error
	if len(path) == 0 {
		ProjectUrl := utils.GetCurrentPath()
		path = ProjectUrl + "./rpc/"
		os.Mkdir(path, os.ModePerm)
		os.Mkdir(path+"/"+arg[0], os.ModePerm)
		os.Mkdir(path+"/"+arg[0]+"/"+arg[1], os.ModePerm)
		protoPath := path + arg[0] + "/" + arg[1] + "/" + arg[0] + "_" + arg[1] + ".proto"
		fmt.Println("初始化proto文件", protoPath)
		protoFile, err = os.Create(protoPath)
	} else {
		fmt.Println("gen创建文件", "./rpc/"+arg[0]+"/"+arg[1]+"/"+arg[0]+"_"+arg[1]+".proto")
		protoFile, err = os.Create("./rpc/" + arg[0] + "/" + arg[1] + "/" + arg[0] + "_" + arg[1] + ".proto")
	}
	RpcRoute.Module = strings.Replace(string(utils.RunCmd("go list -m", true)), "\n", "", -1)
	RpcRoute.UpModule = utils.Camel2Comm(RpcRoute.Module)
	RpcRoute.Package = utils.Lcfirst(arg[0])
	RpcRoute.UpPackage = utils.Ucfirst(arg[0])
	RpcRoute.UpMethod = utils.Ucfirst(arg[1])
	RpcRoute.Method = utils.Lcfirst(arg[1])
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("proto文件生成成功")
	buf := new(bytes.Buffer)
	tmpl, err := template.New("gen-proto").Parse(strings.TrimSpace(tpl.PROTOTPL))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, RpcRoute); err != nil {
		panic(err)
	}
	protoFile.WriteString(buf.String())
	protoFile.Close()
}

func genRpc(path, dir, fileName, method string) {
	fmt.Println("当前rpc路径:" + path)
	fmt.Println("rpc文件名称:" + fileName)
	govalidatorpath := "./internal/api/rpc"
	gorpcpath := "./rpc/"
	//os.Mkdir("./internal/api/http/"+dir, os.ModePerm)
	//os.Mkdir("./internal/api/http/"+dir+"/"+method, os.ModePerm)
	//goginpath := "./internal/api/http/" + dir + "/" + method
	fmt.Println("生成go文件")
	//cmd := "protoc -I " + path + " --proto_path=${GOPATH}/pkg/mod  --proto_path=${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.3.2 --proto_path=. --govalidators_out=" + govalidatorpath + " --go-grpc_out=plugins=grpc:" + path + " --oaago_out=" + path + " --oaago_opt=paths=source_relative " + path + "/" + fileName + ".proto"
	cmd := `protoc -I ./ -I ` + gorpcpath + ` \
               --proto_path=$GOPATH/src \
--proto_path=${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.3.2 \
               --proto_path=${GOPATH}/pkg/mod \
               --proto_path=` + govalidatorpath + ` \
--govalidators_out=paths=source_relative:./internal/api/rpc \
               --go_out=paths=source_relative:./internal/api/rpc \
               --go-grpc_out=` + govalidatorpath + ` --go-grpc_opt=paths=import \
               --oaago_out=` + govalidatorpath + ` \
               --oaago_opt=paths=import \
               --grpc-gateway_out ` + govalidatorpath + ` --grpc-gateway_opt paths=import \
               --grpc-gateway_opt logtostderr=true \
               --grpc-gateway_opt generate_unbound_methods=true \
               --grpc-gateway_opt register_func_suffix=GW \
               --grpc-gateway_opt allow_delete_body=true \
               --doc_out=./docs \
               --doc_opt=html,index.html \
               --openapiv2_out ./docs --openapiv2_opt logtostderr=true \
               ` + path + "/" + fileName + ".proto"
	fmt.Println("生成go文件" + cmd)
	c := exec.Command("bash", "-c", cmd)
	output, err := c.CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		fmt.Println(err.Error())
	}
}
