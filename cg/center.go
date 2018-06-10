package cg

import (
	"strings"
	"github.com/gmaclinuxer/gogame/ipc"
)

// CenterServer must implement ipc.Server interface
var _ ipc.Server = &CenterServer{}

func NewCenterServer() *CenterServer {

	servers := make(map[string]ipc.Server)
	players := make([]*Player, 0)

	return &CenterServer{servers: servers, players: players}
}

func (server *CenterServer) Handle(method, params string) *ipc.Response {

	switch method {
	case "addplayer":
		err := server.addPlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	case "removeplayer":
		err := server.removePlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	case "listplayer":
		players, err := server.listPlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
		return &ipc.Response{"200", players}
	case "broadcast":
		err := server.broadcast(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	default:
		return &ipc.Response{Code: "404", Body: strings.Join([]string{method, params}, ":")}
	}
}

func (server *CenterServer) Name() string {
	return "CenterServer"
}
