package main

import (
	"github.com/Olionnn/mysnvr/common"
	"github.com/Olionnn/mysnvr/web"
)

func main() {

	common.CreateDefaultEnv()

	// go app.StartRecord()
	web.StartWeb()

}
