package main

import (
	_ "test/my-web-app/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

