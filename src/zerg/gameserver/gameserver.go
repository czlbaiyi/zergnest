package gameserver

import (
	"os"
	"os/signal"

	"zerg/log"
	"zerg/module"
)

func StartServer() {
	zerg.Run(
		hero.Module,
		equip.Module,
		room.Module,
	)

	// module
	for i := 0; i < len(mods); i++ {
		module.Register(mods[i])
	}
	module.Init()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Release("Leaf closing down (signal: %v)", sig)
	module.Destroy()
}
