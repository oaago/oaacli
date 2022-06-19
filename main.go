package main

import (
	"fmt"
	"github.com/oaago/oaacli/cmd"
	"github.com/oaago/oaacli/utils"
)

var BaseName = ""

func main() {
	BaseName = utils.GetCurrentPath()
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
	}
}
