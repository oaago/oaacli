package tpl

import _ "embed"

//go:embed api.tpl
var ApiTPL string

//go:embed httpServiceHandler.tpl
var HttpServiceHandler string

//go:embed rpcService.tpl
var RpcServiceTpl string

//go:embed dao.tpl
var DaoTpl string

//go:embed dockerFile.tpl
var DockerFileTpl string

//go:embed httpRouter.tpl
var HttpRouterTpl string

//go:embed rpcRouter.tpl
var RpcRouterTpl string

//go:embed httpServiceTpl.tpl
var HttpServiceTpl string

//go:embed httpTypes.tpl
var HttpTypes string

//go:embed rpcTypes.tpl
var RpcTypesTpl string

//go:embed proto.tpl
var ProtoTpl string

//go:embed middleware.tpl
var MiddlewareTpl string

//go:embed main.tpl
var MainTpl string

//go:embed confing.tpl
var ConfingTpl string

//go:embed powerproto.tpl
var PowerprotoTpl string

//go:embed oaa.tpl
var OAATpl string

//go:embed consts.tpl
var ConstsTpl string
