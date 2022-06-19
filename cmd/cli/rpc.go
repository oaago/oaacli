package cli

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/oaago/cli/cmd/tpl"
	"github.com/oaago/cli/utils"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
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

var GenRpc = &cobra.Command{
	Use:   "rpc-gen",
	Short: "示rpc 中央proto生成器",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("proto gen")
		RpcRoute.Module = strings.Replace(string(utils.RunCmd("go list -m", true)), "\n", "", -1)
		data, _ := os.ReadFile(utils.GetCurrentPath() + "./oaa.json")
		httpData := gjson.Get(string(data), "rpc").Array()
		for _, datum := range httpData {
			s := datum.String()
			if strings.Contains(s, "&/") {
				urlm := strings.Split(s, "&/")
				if len(urlm) == 2 {
					RpcRoute.RPCMap = append(RpcRoute.RPCMap, RpcList{
						Module:    RpcRoute.Module,
						UpModule:  utils.Camel2Comm(RpcRoute.Module),
						UpMethod:  utils.Ucfirst(urlm[1]),
						Method:    urlm[1],
						UpPackage: utils.Ucfirst(urlm[0]),
						Package:   urlm[0],
					})
					arg := strings.Split(urlm[1], "/")
					os.Mkdir("./rpc/"+arg[0], os.ModePerm)
					os.Mkdir("./rpc/"+arg[0]+"/"+arg[1], os.ModePerm)
					ProjectUrl := utils.GetCurrentPath()
					path := ProjectUrl + "rpc/" + arg[0] + "/" + arg[1]
					genProto([]string{urlm[1]}, "./rpc")
					cmd := "protoc -I " + path + " --proto_path=${GOPATH}/pkg/mod  --proto_path=${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.3.2 --govalidators_out=. --go_out=plugins=grpc:" + path + " --go_opt=paths=source_relative --oaago_out " + path + " --oaago_opt=paths=source_relative  " + path + "/" + arg[0] + "_" + arg[1] + ".proto"
					fmt.Println(cmd)
					c := exec.Command("bash", "-c", cmd)
					output, err := c.CombinedOutput()
					fmt.Println(string(output), err)
				} else {
					panic("格式不正确 get&/aa/bb")
				}
			}
		}
	},
}

func genProto(args []string, path string) {
	arg := strings.Split(args[0], "/")
	fmt.Println(arg)
	var protoFile *os.File
	var err error
	if len(path) == 0 {
		ProjectUrl := utils.GetCurrentPath()
		path = ProjectUrl + "internal/api/rpc/"
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
	govalidatorpath := "./internal/api/"
	//os.Mkdir("./internal/api/http/"+dir, os.ModePerm)
	//os.Mkdir("./internal/api/http/"+dir+"/"+method, os.ModePerm)
	//goginpath := "./internal/api/http/" + dir + "/" + method
	cmd := "protoc -I " + path + " --proto_path=${GOPATH}/pkg/mod  --proto_path=${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.3.2 --proto_path=. --govalidators_out=" + govalidatorpath + " --go_out=plugins=grpc:" + path + " --go_opt=paths=source_relative --oaago_out " + path + " --oaago_opt=paths=source_relative " + path + "/" + fileName + ".proto"
	fmt.Println(cmd)
	c := exec.Command("bash", "-c", cmd)
	output, err := c.CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		fmt.Println(err.Error())
	}
}
