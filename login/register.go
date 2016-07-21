package login

import (
	"log"
	"net/rpc/jsonrpc"
	"os"
)

type (
	RegisterRpc struct {
		username string
		passwd string
		regIP string
		regTime string
	}

	RegisterReply struct {

	}
)

func (l *Login)Register(params RegisterRpc, re *RegisterReply) (err error) {
	if len(os.Args) != 2 {
		log.Print("Usage: ", os.Args[0], "server:port")
		log.Fatal(1)
	}
	service := os.Args[1]

	client, err := jsonrpc.Dial("tcp", service)
	defer client.Close()
	if err != nil {
		return err
	}

	args := params
	err = client.Call("Db.UserRegister", args, &re)
	if err != nil {
		return err
	}

	return
}
