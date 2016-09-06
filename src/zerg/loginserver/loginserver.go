package loginserver

import (
	"io/ioutil"
	"net/http"

	"zerg/log"

	PB_Login "zerg/protobuf/PB_Login/go"

	"zerg/conf"

	"github.com/golang/protobuf/proto"
)

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
	newLoginRet := &PB_Login.CS11001{}
	err = proto.Unmarshal(body, newLoginRet)
	if err != nil {
		log.Error("unmarshaling error: ", err)
	} else {
		log.Debug(string(body))
	}

}

func StartServer() {
	http.HandleFunc("/", serveHome)
	err := http.ListenAndServe(conf.LoginServerPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
