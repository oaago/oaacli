package middleware

import (
    "fmt"
    "github.com/oaago/server/v2/types"
    "github.com/oaago/server/v2/http/core"
)

//Pid 局部中间件
type Pid types.PartMiddleware

//Gid 全局中间件
type Gid types.GlobalMiddleware

//NewPid 局部示例
func NewPid() Pid {
	return Pid{}
}
func (Pid) TT(c *types.Context) {
	fmt.Println("PartMiddleware")
	c.Next()
	fmt.Println("2222")
}

//NewGid 全局中间件示例
func NewGid() Gid {
	return Gid{}
}

func (Gid) BB(c *types.Context) {
	fmt.Println("GlobalMiddleware")
	c.Next()
	fmt.Println("4444")
}