package main

import (
	//命令行选项解析器
	"flag"
	"fmt"
	"os"

	"zerg/conf"
	"zerg/dataserver"
	"zerg/gateserver"
	"zerg/log"
	"zerg/loginserver"
	"zerg/mailserver"
	"zerg/rankserver"
	"zerg/sonyflake"
	"zerg/statusserver"
)

var serverName = ""

//参数检查
func argsInit() {
	flag.Parse()
	argsNum := flag.NArg()
	if argsNum != 1 {
		log.Debug("please set args, the args will start some server")
		os.Exit(-1)
		return
	}
	serverName = flag.Arg(0)
}

//配置检查
func configCheck() {
	var sf *sonyflake.Sonyflake
	var st sonyflake.Settings

	st.MachineID = func() (uint16, error) {
		if conf.MACHINE_ID < 0 {
			fmt.Errorf("no machine id")
		}
		return conf.MACHINE_ID, nil
	}

	sf = sonyflake.NewSonyflake(st)

	if sf == nil {
		panic("sonyflake not created")
	}
}

//启动对应服务器
func startServer() {
	switch serverName {
	case "loginserver":
		loginserver.StartServer()
		log.Debug("begin start login server")
	case "statusserver":
		statusserver.StartServer()
		log.Debug("begin start data server")
	case "gateserver":
		gateserver.StartServer()
		log.Debug("begin start gate server")
	case "dataserver":
		dataserver.StartServer()
		log.Debug("begin start data server")
	case "rankserver":
		rankserver.StartServer()
		log.Debug("begin start rank server")
	case "mailserver":
		mailserver.StartServer()
		log.Debug("begin start mail server")
	default:
		log.Fatal("star serverName error, the server server named :" + serverName + "has't found")
		os.Exit(-1)
	}
}

func main() {
	// logger
	if conf.LogLevel != "" {
		logger, err := log.New(conf.LogLevel, conf.LogPath)
		if err != nil {
			panic(err)
		}
		log.Export(logger)
		defer logger.Close()
	}

	log.Release("Leaf %v starting up", version)

	argsInit()

	configCheck()

	startServer()
}
