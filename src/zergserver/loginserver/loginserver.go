package loginserver

import (
	"io/ioutil"
	"net/http"

	"zerg/log"
	"zerg/zerror"

	"zerg/conf"
	"zerg/routermanager"
)

type LoginServer struct {
}

func (moudle *LoginServer) Load() {
	RouterManageInit()
}

func (moudle *LoginServer) Start() {
	http.HandleFunc("/", serveHome)
	url := conf.GetCurrentLoginServerSeparatorPort()
	log.Release("Start Server at url :" + url)
	err := http.ListenAndServe(url, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (moudle *LoginServer) Destory() {

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(zerror.New("http:", err.Error()).Error()))
		return
	} else {
		retBuf, err := routermanager.DoMessageBuffs(body)
		if err != nil {
			w.Write([]byte(zerror.New("http:", err.Error()).Error()))
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("content-type", "text/plain")
			w.Write(retBuf)
		}
	}
}

func RouterManageInit() {
	var router routermanager.Router
	router = new(LoginRouter)
	routermanager.RegisterMessageRouter(&router)
}
