package main

import (
	"fmt"
	"log"
	"net"
	"strings"
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
	c.name = args[1]
	c.msg(fmt.Sprintf("name changed to : %s", c.name))
}

func (s *server) join(c *client, args []string) {
	targetName := args[1]

	r, exist := s.rooms[targetName]
	if !exist {
		r = &room{
			name:    targetName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[targetName] = r
	}

	r.members[c.connection.RemoteAddr()] = c

	s.quitCurrentRoom(c)

	c.room = r

	r.broadcast(c, fmt.Sprintf("%s has joined the chat", c.name))
}

func (s *server) listRooms(c *client, args []string) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}

	c.msg(fmt.Sprintf("available rooms are : %s", strings.Join(rooms, ", ")))
}

func (s *server) msg(c *client, args []string) {

}

func (s *server) quit(c *client, args []string) {

}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		delete(c.room.members, c.connection.RemoteAddr())
		c.room.broadcast(c, fmt.Sprintf("%s has left the chat", c.name))
	}
}
