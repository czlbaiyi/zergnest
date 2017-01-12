package gameserver

import (
	"zerg/routermanager"
	"zerg/zerror"

	"zergserver/gameserver/protobuf/go"

	"github.com/golang/protobuf/proto"
)

type GameRouter struct {
}

func (r *GameRouter) GetRouterType() int32 {
	return int32(gamemsg.CS21001_PB_PackType)
}

func (r *GameRouter) OnPacketHandler(bytes []byte) ([]byte, error) {
	router := &gamemsg.CS21001{}
	err := proto.Unmarshal(bytes, router)
	if err != nil {
		return nil, zerror.New("GameRouter：", err.Error())
	}

	retRouter := &gamemsg.SC21002{}
	dates, err := proto.Marshal(retRouter)
	if err != nil {
		return nil, zerror.New("GameRouter：", err.Error())
	}
	return routermanager.GetMsgPacket(1, int32(gamemsg.SC21002_PB_PackType), dates)
}
