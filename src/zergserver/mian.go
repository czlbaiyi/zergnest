package main

import (
	"fmt"
	"os"
	"runtime"

	"strconv"
	"zerg/conf"
	"zerg/log"
	"zerg/moudle"
	"zerg/sonyflake"
	"zergserver/gameserver"
	"zergserver/loginserver"
)

var curMoudle moudle.Moudle

func loadConfig() {
	conf.ReadConfig()
}

//配置检查
func configCheck() {
	var sf *sonyflake.Sonyflake
	var st sonyflake.Settings

	st.MachineID = func() (uint16, error) {
		if conf.Config.MachineId < 0 {
			fmt.Errorf("no machine id")
		}
		return conf.Config.MachineId, nil
	}

	sf = sonyflake.NewSonyflake(st)

	if sf == nil {
		panic("sonyflake not created")
	}
}

//启动对应服务器
func startServer() {
	switch conf.Config.StartType {
	case 0:
		curMoudle = new(loginserver.LoginServer)
		log.Debug("begin start LoginServer")
	case 1:
		curMoudle = new(gameserver.GameServer)
		log.Debug("begin start GameServer")
	default:
		log.Fatal("star server error, the server type was not in list, type = " + strconv.Itoa(conf.Config.StartType))
		os.Exit(-1)
	}
	curMoudle.Load()
	curMoudle.Start()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	loadConfig()
	configCheck()
	log.Init(conf.Config.LogLevel, conf.Config.LogPath)
	log.Release("zergnest start")

	startServer()
	i := 0
	for a := 0; a < 100000000000; a++ {
		i = i + a*2
	}
	log.Close()
	log.Release("zergnest close")
	for a := 0; a < 100000000000; a++ {
		i = i + a*2
	}
	log.Release(strconv.Itoa(i))
}
