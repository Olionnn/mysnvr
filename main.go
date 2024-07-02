package main

import (
	"github.com/Olionnn/mysnvr/app"
	"github.com/Olionnn/mysnvr/web"
)

func main() {

	go app.StartRecord()
	web.StartWeb()

}
