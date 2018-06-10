package cg

import (
	"encoding/json"
	"errors"
)

func (s *CenterServer) addPlayer(params string) error {

	player := NewPlayer()

	err := json.Unmarshal([]byte(params), &player)
	if err != nil {
		return err
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	// todo duplicate login check
	s.players = append(s.players, player)

	return nil
}

func (s *CenterServer) removePlayer(params string) error {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i, v := range s.players {
		if v.Name == params {
			s.players = append(s.players[:i], s.players[i+1:]...)
			return nil
		}
	}

	return errors.New("player not found")
}

func (s *CenterServer) listPlayer(params string) (players string, err error) {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.players) > 0 {
		b, _ := json.Marshal(s.players)
		players = string(b)
	} else {
		err = errors.New("no player online")
	}
	return
}

func (s *CenterServer) broadcast(params string) error {

	var msg Message
	if err := json.Unmarshal([]byte(params), &msg); err != nil {
		return err
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.players) > 0 {
		for _, p := range s.players {
			p.mq <- &msg
		}
		return nil
	} else {
		return errors.New("no player online")
	}

}
