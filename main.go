package main

import (
	"log"
	"net"
)

func main() {
	serv := InitServer()
	go serv.run()

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("Unable to start Server : %s", err.Error())
	}

	defer listener.Close()
	log.Print("started serv on port :8888")

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Printf("unable to accept connection : %s", err.Error())
		}

		go serv.newClient(connection)
	}
}
