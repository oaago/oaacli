package main

import (
	"fmt"

	"github.com/oaago/cli/cmd"
	"github.com/oaago/cli/utils"
)

var BaseName = ""

func main() {
	BaseName = utils.GetCurrentPath()
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
	}
}
