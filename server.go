package main

import (
	"log"
	"net"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func InitServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listRooms(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client, cmd.args)
		}
	}
}

func (s *server) newClient(connection net.Conn) {
	log.Printf("new clien has connected : %s", connection.RemoteAddr().String())

	client := &client{
		connection: connection,
		name:       "anonymous",
		commands:   s.commands,
	}

	client.readInput()
}

func (s *server) nick(c *client, args []string) {

}

func (s *server) join(c *client, args []string) {

}

func (s *server) listRooms(c *client, args []string) {

}

func (s *server) msg(c *client, args []string) {

}

func (s *server) quit(c *client, args []string) {

}
