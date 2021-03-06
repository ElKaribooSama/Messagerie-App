package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	connection net.Conn
	name       string
	room       *room
	commands   chan<- command
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.connection).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/nick":
			c.commands <- command{
				id:     CMD_NICK,
				client: c,
				args:   args,
			}
		case "/join":
			c.commands <- command{
				id:     CMD_JOIN,
				client: c,
				args:   args,
			}
		case "/rooms":
			c.commands <- command{
				id:     CMD_ROOMS,
				client: c,
				args:   args,
			}
		case "/msg":
			c.commands <- command{
				id:     CMD_MSG,
				client: c,
				args:   args,
			}
		case "/quit":
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
				args:   args,
			}
		default:
			c.err(fmt.Errorf("unknow command: %s", cmd))
		}
	}
}

func (c *client) err(err error) {
	c.connection.Write([]byte("Error: " + err.Error() + "\n"))
}

func (c *client) msg(msg string) {
	c.connection.Write([]byte("> " + msg + "\n"))
}
