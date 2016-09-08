package loginservermsg

import "zerg/routermanager"

func Init() {
	var router routermanager.Router
	var routerL LoginRouter
	router = routerL
	routermanager.RegisterMessageRouter(&router)
}
