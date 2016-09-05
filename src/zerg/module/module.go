package module

import (
	"os"
	"os/signal"
	"runtime"
	"sync"

	"zerg/conf"
	"zerg/log"
)

type Module interface {
	OnInit()
	OnDestroy()
	Run(closeSig chan bool)
}

type module struct {
	mi       Module
	closeSig chan bool
	wg       sync.WaitGroup
}

var mods []*module

func Register(mi Module) {
	m := new(module)
	m.mi = mi
	m.closeSig = make(chan bool, 1)

	mods = append(mods, m)
}

func Init() {
	for i := 0; i < len(mods); i++ {
		mods[i].mi.OnInit()
	}

	for i := 0; i < len(mods); i++ {
		m := mods[i]
		m.wg.Add(1)
		go run(m)
	}
}

func Destroy() {
	for i := len(mods) - 1; i >= 0; i-- {
		m := mods[i]
		m.closeSig <- true
		m.wg.Wait()
		destroy(m)
	}
}

func run(m *module) {
	m.mi.Run(m.closeSig)
	m.wg.Done()
}

func destroy(m *module) {
	defer func() {
		if r := recover(); r != nil {
			if conf.LenStackBuf > 0 {
				buf := make([]byte, conf.LenStackBuf)
				l := runtime.Stack(buf, false)
				log.Error("%v: %s", r, buf[:l])
			} else {
				log.Error("%v", r)
			}
		}
	}()

	m.mi.OnDestroy()
}

func Run(mods ...Module) {
	// logger
	if conf.LogLevel != "" {
		logger, err := log.New(conf.LogLevel, conf.LogPath)
		if err != nil {
			panic(err)
		}
		log.Export(logger)
		defer logger.Close()
	}

	log.Release("Leaf starting up")

	// module
	for i := 0; i < len(mods); i++ {
		Register(mods[i])
	}
	Init()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Release("Leaf closing down (signal: %v)", sig)
	Destroy()
}
