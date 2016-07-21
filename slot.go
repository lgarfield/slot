package main

import (
	// module
	"slot/db"
	"slot/login"

	// Internal
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"sync"
)

type (

)

func main() {
	login := new(login.Login)
	rpc.Register(login)

	dbRpc := new(db.Db)
	rpc.Register(dbRpc)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	defer listener.Close()
	checkError(err)

	mutex := &sync.Mutex{}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Fatal error:", err.Error())

			continue
		}

		go func () {
			mutex.Lock()

			jsonrpc.ServeConn(conn)

			mutex.Unlock()
		}()
	}
}

func checkError(err error) {
	if err != nil {
		log.Println("Fatal error:", err.Error())
		os.Exit(1)
	}
}
