package main

import (
	"github.com/Olionnn/mysnvr/app"
	"github.com/Olionnn/mysnvr/web"
)

func main() {
	fmt.Println("Hello, World!")
	go app.StartRecord()
	web.StartWeb()

}
