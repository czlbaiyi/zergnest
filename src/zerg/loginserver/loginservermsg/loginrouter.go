package loginservermsg

import (
	"zerg/log"
	"zerg/protobuf/go"
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
	} else {
		log.Debug("return some things")
		return nil, nil
	}
}
