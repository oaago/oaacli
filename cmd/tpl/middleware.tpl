package middleware

import (
    "fmt"
    "github.com/oaago/server/oaa"
)

// 局部中间件
type Pid oaa.PartMiddleware

// 全局中间件
type Gid oaa.GlobalMiddleware

// 局部示例
func NewPid() Pid {
	return Pid{}
}
func (Pid) TT(c *oaa.Ctx) {
	fmt.Println("PartMiddleware")
	c.Next()
	fmt.Println("2222")
}

// 全局中间件示例
func NewGid() Gid {
	return Gid{}
}

func (Gid) BB(c*oaa.Ctx) {
	fmt.Println("GlobalMiddleware")
	c.Next()
	fmt.Println("4444")
}