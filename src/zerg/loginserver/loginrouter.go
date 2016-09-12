package loginserver

import (
	"zerg/conf"
	"zerg/protobuf/go"
	"zerg/routermanager"
	"zerg/zerror"

	"github.com/golang/protobuf/proto"
)

type LoginRouter struct {
}

func (r LoginRouter) GetRouterType() int32 {
	return int32(pb_message.CS11001_PB_PackType)
}

func (r LoginRouter) OnPacketHandler(bytes []byte) ([]byte, error) {
	router := &pb_message.CS11001{}
	err := proto.Unmarshal(bytes, router)
	if err != nil {
		return nil, zerror.New("LoginRouter：", err.Error())
	}

	loginRet, errStr := LoginOnServer(router.PlatformAccountId, router.AuthId)
	if loginRet == true {
		retRouter := &pb_message.SC11002{}
		retRouter.SessionId = "11111111111"
		retRouter.LocalPlayerId = 22222222222
		retRouter.PlatformAccountId = "accid"
		retRouter.PlatformAccountName = "name"
		retRouter.ServerInfos = append(retRouter.ServerInfos, &pb_message.PB_ServerInfo{
			1,
			"TestServer",
			conf.LoginServerUrls[0].GetServerUrl(),
			pb_message.PB_ServerInfo_good,
		})

		dates, err := proto.Marshal(retRouter)
		if err != nil {
			return nil, zerror.New("LoginRouter：", err.Error())
		}
		return routermanager.GetMsgPacket(1, int32(pb_message.SC11002_PB_PackType), dates)
	} else {
		return nil, zerror.New("LoginRouter：", errStr)
	}
}
