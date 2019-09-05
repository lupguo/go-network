package main

import (
	"encoding/gob"
	"fmt"
	"github.com/tkstorm/go-network/tcp/internal/helper"
	"log"
	"net"
)

const listenAddr = "0.0.0.0:6666"

func main() {
	// net listen
	ln, err := net.Listen("tcp4", listenAddr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("tcp server listen on", listenAddr)
	for {
		// conn accept
		conn, err := ln.Accept()
		if err != nil {
			log.Println("conn accept error", err)
			continue
		}
		// config setting
		//_ = conn.SetDeadline(time.Now().Add(15 * time.Second))

		// read from conn
		go HandleAsciiConn(conn)
	}

}

//HandleAsciiConn
func HandleAsciiConn(conn net.Conn) {
	defer func() {
		if p := recover(); p != nil {
			log.Println(p)
		}
	}()
	defer conn.Close()
	msg, err := helper.ReadConnMessage(conn)
	if err != nil {
		log.Println(err)
	}
	helper.Prompt(conn.RemoteAddr().String(), msg)
	fmt.Fprintln(conn, "got")
}

//HandleGobConn
func HandleGobConn(conn net.Conn) {
	defer conn.Close()

	// receive binary data from client conn & handle
	var msg string
	if err := gob.NewDecoder(conn).Decode(&msg); err != nil {
		log.Println("server decode err,", err)
	}
	log.Println("server receive > ", msg)

	// send data to client conn
	var replay = "server finish!"
	if err := gob.NewEncoder(conn).Encode(replay); err != nil {
		log.Println("server gob encode err", err)
	}
}
