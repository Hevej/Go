package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

// Listen takes a network type (TCP) and an addresss
// and por to bind, and returns a net.Listener

/*
type Listener interface {
	// Accept waits for and returns the next connection to the listener
	Accept() (c Conn, err error)

	// Close closes the listener
	// Any blocked Accept operations will be unblocked and return errors
	Close() error
	Addr() Addr
}*/

func server() {
	// listen on a port
	// Mediante TCP al puerto 9999
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	_ = c.Close()
}

func client() {
	// connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	// send the message
	msg := "Hello, World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
