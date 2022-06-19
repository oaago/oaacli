package main

import (
	"fmt"

	"github.com/oaago/oaago/utils"
	"github.com/oaago/oaago/cmd"
)

var BaseName = ""

func main() {
	BaseName = utils.GetCurrentPath()
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
	}
}
