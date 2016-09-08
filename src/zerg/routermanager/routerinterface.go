package routermanager

type Router interface {
	GetRouterType() int32
	OnPacketHandler([]byte) ([]byte, error)
}
