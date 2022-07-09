package main

import (
	"fmt"
	"time"

	"github.com/oaago/oaago/cmd"
	"github.com/oaago/oaago/utils"
)

var BaseName = ""

func main() {
	BaseName = utils.GetCurrentPath()
	t := time.Now()
	fmt.Println("当前时间 " + t.Format("2006-01-02 15:04:05"))
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
	}
}
