package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type serverUrl struct {
	Ip   string
	Port string
}

type config struct {
	StartTypeDes   string
	StartType      int
	StartIdxDes    string
	StartIdx       int
	LoginServerUrl []*serverUrl
	GameServerUrls []*serverUrl
	logDes         string
	LogLevel       int
	LogPath        string
	LenStackBuf    int
	MachineId      uint16
}

var Config config

func (s *serverUrl) GetServerUrl() string {
	var url string = s.Ip
	url += ":"
	url += s.Port
	return url
}

func (s *serverUrl) GetServerSeparatorPort() string {
	var url string = ":"
	url += s.Port
	return url
}

func GetCurrentGameServerUrl() string {
	var s *serverUrl = Config.GameServerUrls[Config.StartIdx]
	var url string = s.Ip
	url += ":"
	url += s.Port
	return url
}

func GetCurrentLoginServerSeparatorPort() string {
	var s *serverUrl = Config.LoginServerUrl[Config.StartIdx]
	var url string = ":"
	url += s.Port
	return url
}

func ReadConfig() {
	fi, err := os.Open("res/config.json")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(fd))

	err = json.Unmarshal(fd, &Config)
	if err != nil {
		panic(err)
	}

	fmt.Println(Config)

}
