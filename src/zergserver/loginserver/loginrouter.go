package loginserver

import (
	"zerg/conf"
	"zerg/routermanager"
	"zerg/zerror"
	"zergserver/loginserver/protobuf/go"

	"github.com/golang/protobuf/proto"
)

type LoginRouter struct {
}

func (r *LoginRouter) GetRouterType() int32 {
	return int32(loginmsg.CS11001_PB_PackType)
}

func (r *LoginRouter) OnPacketHandler(bytes []byte) ([]byte, error) {
	router := &loginmsg.CS11001{}
	err := proto.Unmarshal(bytes, router)
	if err != nil {
		return nil, zerror.New("LoginRouter：", err.Error())
	}

	loginRet, errStr := login(router.Account, router.Password)
	if loginRet == true {
		retRouter := &loginmsg.SC11002{}
		retRouter.SessionId = "11111111111"
		retRouter.PlayerId = 2222222222
		for index := 0; index < len(conf.Config.GameServerUrls); index++ {
			retRouter.ServerInfos = append(retRouter.ServerInfos, &loginmsg.PB_ServerInfo{
				1,
				"TestServer",
				conf.Config.GameServerUrls[index].GetServerUrl(),
				loginmsg.PB_ServerInfo_good,
			})
		}

		dates, err := proto.Marshal(retRouter)
		if err != nil {
			return nil, zerror.New("LoginRouter：", err.Error())
		}
		return routermanager.GetMsgPacket(1, int32(loginmsg.SC11002_PB_PackType), dates)
	} else {
		return nil, zerror.New("LoginRouter：", errStr)
	}
}

func login(accout string, password string) (bool, string) {
	if accout == "" {
		return false, "accout must not empty"
	}

	if password == "" {
		return false, "password must not empty"
	}

	if playerData, ok := playersData[accout]; !ok {
		return addPlayer(accout, password)
	} else if playerData == password {
		return true, ""
	} else {
		return false, "password errer"
	}
}
