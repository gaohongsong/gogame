package cg

import "fmt"

func NewPlayer() *Player {

	// Message Channel with buffer
	mq := make(chan *Message, 1024)
	player := &Player{mq: mq}

	// Start a new routine to serve new player
	go func(p *Player) {
		for {
			msg := <-player.mq
			fmt.Printf("%s received message: %s", p.Name, msg.Content)
		}
	}(player)

	return player
}
