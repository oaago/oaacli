package middleware

import (
    "fmt"
    h "github.com/oaago/server/v2/http"
)

//Pid 局部中间件
type Pid h.PartMiddleware

//Gid 全局中间件
type Gid h.GlobalMiddleware

//NewPid 局部示例
func NewPid() Pid {
	return Pid{}
}
func (Pid) TT(c *h.Context) {
	fmt.Println("PartMiddleware")
	c.Next()
	fmt.Println("2222")
}

//NewGid 全局中间件示例
func NewGid() Gid {
	return Gid{}
}

func (Gid) BB(c*h.Context) {
	fmt.Println("GlobalMiddleware")
	c.Next()
	fmt.Println("4444")
}