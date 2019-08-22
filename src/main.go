package main

import (
	"log"
	"mtblog/dao"
	_ "mtblog/routers"
	"github.com/astaxie/beego"
)

func main() {
	log.Print(dao.DB)
	beego.Run()
}

