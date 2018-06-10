package ipc

import "encoding/json"

type IpcClient struct {
	con Channel
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()

	return &IpcClient{c}
}

func (cli *IpcClient) Call(method, params string) (resp *Response, err error) {

	req := &Request{method, params}

	var b []byte

	b, err = json.Marshal(req)
	if err != nil {
		return
	}

	// send to server by chan
	cli.con <- string(b)

	// get resp from server by chan
	str := <-cli.con

	// todo why not just user resp for unmarshal
	var res Response
	err = json.Unmarshal([]byte(str), &res)
	resp = &res

	return
}

func (cli *IpcClient) Close() {
	cli.con <- "CLOSE"
}
