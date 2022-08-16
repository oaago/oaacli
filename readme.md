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
   
    go install github.com/oaago/oaago@main
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
   init        oaacli init 根据 defined.json 生成出来需要的项目文件，可以制定配置文件oaa.json
   new         示例 oaacli new project 生成项目 包含了http+rpc
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

3 .oaa.json 的定义示例

   ```
{
  "http": [
    "get,post,delete,put@/user/tenant**企业租户",
    "get,post,delete,put@/user/tenant/role**企业租户-用户-角色",
    "*@/sys/permission**用户权限",
    "*@/sys/permission/group**用户权限组",
    "*@/sys/role**用户角色",
    "*@/sys/role/group**用户角色组",
    "*@/sys/dict**系统字典",
    "*@/sys/dict/item**系统字典明细",
    "*@/sys/user**用户表"
    ]
}
   ```

3.1
get,post,delete,put 代表的是请求方式 * ,代表支持所有的请求方式(get, post, delete, put, patch)
/user/tenant 代表的是数据表 user_tenant
**企业租户 **后面代表的是备注

所有api/rpc的生成都依赖于 internal/defined.json 的规则