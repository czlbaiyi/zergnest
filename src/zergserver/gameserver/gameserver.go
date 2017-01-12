package gameserver

import (
	"net/http"
	"zerg/conf"
	"zerg/log"

	"zerg/routermanager"

	"golang.org/x/net/websocket"
)

type GameServer struct {
}

func (moudle *GameServer) Load() {
	RouterManageInit()
}

func (moudle *GameServer) Start() {
	go h.run()

	http.Handle("/", websocket.Handler(onWebSocket))
	// 开始服务
	url := conf.GetCurrentGameServerUrl()
	log.Release("Start Server at url :" + url)
	err := http.ListenAndServe(url, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (moudle *GameServer) Destory() {

}

func onWebSocket(ws *websocket.Conn) {
	c := &connection{send: make(chan []byte, 256), ws: ws}
	h.register <- c
	defer func() { h.unregister <- c }()
	go c.writer()
	c.reader()
}

func RouterManageInit() {
	var router routermanager.Router
	router = new(GameRouter)
	routermanager.RegisterMessageRouter(&router)
}
