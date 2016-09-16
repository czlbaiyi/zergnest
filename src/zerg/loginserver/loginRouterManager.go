package loginserver

import "zerg/routermanager"

func RouterManageInit() {
	var router routermanager.Router
	var routerL LoginRouter
	router = routerL
	routermanager.RegisterMessageRouter(&router)
}
