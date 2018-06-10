package cg

import (
	"encoding/json"
	"errors"
	"github.com/gmaclinuxer/gogame/ipc"
)

type CenterClient struct {
	*ipc.IpcClient
}

func (cli *CenterClient) AddPlayer(player *Player) error {

	if b, err := json.Marshal(*player); err != nil {
		return err
	} else {
		res, err := cli.Call("addplayer", string(b))
		if err == nil && res.Code == "200" {
			return nil
		}
		return err
	}

}

func (cli *CenterClient) RemovePlayer(name string) error {
	res, err := cli.Call("removeplayer", name)
	if err == nil && res.Code == "200" {
		return nil
	}
	return errors.New(res.Code)

}

func (cli *CenterClient) ListPlayer(params string) (ps []*Player, err error) {

	res, _ := cli.Call("listplayer", params)
	if res.Code != "200" {
		err = errors.New(res.Code)
		return
	}
	err = json.Unmarshal([]byte(res.Body), &ps)

	return
}

func (cli *CenterClient) BroadCast(message string) error {

	msg, err := json.Marshal(&Message{Content: message})

	if err != nil {
		return err
	}

	res, _ := cli.Call("broadcast", string(msg))
	if res.Code != "200" {
		return errors.New(res.Code)
	}

	return nil
}
