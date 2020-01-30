package app

import (
	"blog/app/router"
	"blog/app/dao"
)

func Run() {
	dao.Init()
	router.RegisteRouter()
}
