package cg

import (
	"sync"
	"github.com/gmaclinuxer/gogame/ipc"
)

type Player struct {
	Name  string
	Level int
	Exp   int
	Room  int

	mq chan *Message
}

type Room struct {
	No   int    `json:"no"`
	Name string `json:"name"`
	Cap  int    `json:"cap"`
	Size int    `json:"size"`
}

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

type CenterServer struct {
	servers map[string]ipc.Server
	players []*Player
	rooms   []*Room
	mutex   sync.RWMutex
}
