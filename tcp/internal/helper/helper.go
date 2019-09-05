package helper

import (
	"errors"
	"fmt"
	"log"
	"net"
)

//Prompt conn remote address
func Prompt(prefix string, msg *Message) {
	log.Printf("# %s send [%d]bytes data > %q\n", prefix, msg.len, msg.content)
}

type Message struct {
	len     int
	content []byte
}

func NewMessage(len int, content []byte) *Message {
	return &Message{len: len, content: content}
}

//ReadConnMessage return tcp message in form Message or error when catch any error
func ReadConnMessage(conn net.Conn) (msg *Message, err error) {
	if conn == nil {
		return nil, errors.New("connection is nil")
	}
	// message scan
	msg = new(Message)
	if _, err = fmt.Fscanln(conn, &msg.content); err != nil {
		msg.len = len(msg.content)
		return nil, err
	}
	return
}
