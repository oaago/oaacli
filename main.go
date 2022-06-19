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
	fmt.Println(time.Now().Unix())
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
	}
}
