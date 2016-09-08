package routermanager

import (
	"zerg/protobuf/go"
	"zerg/zerror"

	"github.com/golang/protobuf/proto"
)

var msgMap = make(map[int32]*Router)

func DoMessageBuffs(buffs []byte) ([]byte, error) {
	msg := &pb_message.PB_CommonMsg{}
	err := proto.Unmarshal(buffs, msg)
	if err != nil {
		return nil, zerror.New("protobuf,", err.Error())
	}

	if _, ok := msgMap[msg.OpCode]; !ok {
		return nil, zerror.New("protobuf,", err.Error())
	}

	retBuffs, err := (*msgMap[msg.OpCode]).OnPacketHandler(msg.MsgBuf)
	if err != nil {
		errBuf, err := RetErrorMsg(pb_message.SC10000_PB_MsgDealError, msg.OpCode, err.Error())
		if err != nil {
			return nil, zerror.New("protobuf,", err.Error())
		} else {
			return errBuf, nil
		}

	}
	return retBuffs, nil
}

func RegisterMessageRouter(r *Router) {
	redId := (*r).GetRouterType()
	msgMap[redId] = r
}

func RetErrorMsg(errorCode pb_message.SC10000ErrorCodeType, errOpcode int32, errMsg string) ([]byte, error) {
	router := &pb_message.SC10000{}
	router.ErrorCode = errorCode
	router.CsOpcode = errOpcode
	router.ErrorMsg = errMsg
	routerBytes, err := proto.Marshal(router)
	if err != nil {
		return nil, zerror.New("protobuf,", err.Error())
	}

	msg := &pb_message.PB_CommonMsg{}
	msg.MsgId = 0
	msg.OpCode = errOpcode
	msg.MsgBuf = routerBytes
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return nil, zerror.New("protobuf,", err.Error())
	}

	return msgBytes, nil
}
