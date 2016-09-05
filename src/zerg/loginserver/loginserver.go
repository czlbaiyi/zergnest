package loginserver

import (
	"net/http"

	"zerg/log"

	PB_Login "zerg/protobuf/PB_Login"

	"zerg/conf"
	"github.com/golang/protobuf/proto"
)

func serveHome(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	/*
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("Hello World!"))
		} else {
			temp := string(body)
			temp += "123"
			w.Write([]byte(temp))
		}
	*/

	LoginRet := &PB_Login.SC11002{
		// 使用辅助函数设置域的值
		PlatformAccountId: "hello",
	}

	data, err := proto.Marshal(LoginRet)
	if err != nil {
		log.Error("marshaling error: ", err)
	}

	newLoginRet := &PB_Login.SC11002{}
	err = proto.Unmarshal(data, newLoginRet)
	if err != nil {
		log.Error("unmarshaling error: ", err)
	}

}

func StartServer() {
	http.HandleFunc("/", serveHome)
	err := http.ListenAndServe(conf.LoginServerPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
