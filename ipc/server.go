package ipc

import (
	"encoding/json"
	"fmt"
)

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

// struct must ensure to implement embedded interface's methods
type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer) Connect() Channel {
	session := make(Channel, 0)

	go func(c Channel) {
		for {
			// read from client
			rc := <-c

			// exit case
			if rc == "CLOSE" {
				break
			}

			var req Request
			err := json.Unmarshal([]byte(rc), &req)
			if err != nil {
				fmt.Println("invalid request format: ", req)
				return
			}

			// call real handler
			res := server.Handle(req.Method, req.Params)

			// send response to chan
			wc, err := json.Marshal(res)
			if err != nil {
				fmt.Println("invalid response format: ", res)
				return
			}

			c <- string(wc)

		}
	}(session)

	fmt.Println("a new session has been created successfully.")

	return session
}
