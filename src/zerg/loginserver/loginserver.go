package loginserver

import (
	"io/ioutil"
	"net/http"

	"zerg/log"
	"zerg/zerror"

	"zerg/conf"

	"zerg/loginserver/loginservermsg"
	"zerg/routermanager"
)

func initMondle() {
	loginservermsg.Init()
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
			w.Write(retBuf)
		}
	}
}

func StartServer() {
	initMondle()
	http.HandleFunc("/", serveHome)
	err := http.ListenAndServe(conf.LoginServerPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
