# Oaa CLI TOOL

## 起步

```
请确保安装了以下依赖(请使用1.18以上版本):

- [go 1.18](https://golang.org/dl/)
- [protoc](https://github.com/protocolbuffers/protobuf)
- [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go)
```

1. 使用(目前开发环境支持mac)
    ```
    brew install protobuf
   
    go install github.com/oaago/oaacli@main
    go install github.com/oaago/protoc-gen-oaago@main
    go install github.com/golang/protobuf/protoc-gen-go@v1.5.2
    go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators@latest
    ```

## 2. 介绍

1. 查看所有命令详情
   ```
   oaacli help
   
   Usage:
   oaacli [command]
   
   Available Commands:
   init        oaacli init 根据 defined.json 生成出来需要的项目文件， 可以制定配置文件defined.json
   api         示例 oaacli api get@/staff/info 生成一个api + service
   rpc-gen     直接根据 配置文件defined.json 来生成rpc的 proto 并且 编译好 go文件
   rpc-add     直接根据 oaacli api staff/info 来生成rpc的 proto 并且 编译好 go文件
   api         示例 oaacli api get@/staff/info 生成一个api + service
   completion  Generate the autocompletion script for the specified shell
   help        Help about any command
   model       Print the version number of oaacli model(mysql)
   new         示例 oaacli new project 生成项目 包含了http+rpc
   srv         oaacli service name 根据name 生成service
   types       oaacli types 生成文件
   v           oaacli version 更新时间/更新版本
   Flags:
   -h, --help   help for oaacli
   
   Use "oaacli [command] --help" for more information about a command.
   ```

2. 生成新项目
   ```
   oaacli new oaagotpl
   cd oaagotpl
   oaacli init
   go mod tidy
   go run main.go
   ```

3 .defined 的定义示例

   ```
   {
     "app": [ //app是目录
       "get@/app/bbb"， // app前面的 / 不可省略 bbb代表后续生成的方法 get代表接口请求方式 @代表生成的是http请求
       "get&/aa/bbb"， // app前面的 / 不可省略 bbb代表后续生成的方法  @代表生成的是rpc请求
     ]
   }
   ```

所有api/rpc的生成都依赖于 internal/defined.json 的规则

3 .api的生成(同时会生成service)

   ```
   oaacli api get@/app/aaa
   ```

4 .rpc的生成

   ```
   oaacli rpc-add get@app/aaa
   ```

5 .rpc的生成

   ```
   oaacli srv app/aaa
   ```
