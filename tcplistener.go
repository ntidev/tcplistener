package main

import (
	"fmt"
	"net"
	"strconv"
	"os"
	"log"
	"os/signal"
)

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)
	go func() {
		s := <-sigs
		log.Printf("RECEIVED SIGNAL: %s",s)
		os.Exit(1)
	}()

    args := os.Args
    intf := "0.0.0.0"
    port := 5001

    if len(args) > 1 {
        intf = args[1]
    }

    if len(args) > 2 {
        port, _ = strconv.Atoi(args[2])
    }

    ln, err := net.Listen("tcp", intf + ":" + strconv.Itoa(port))
    if err != nil {
		log.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Listening for connections on %s:%d\n", intf, port)

    for {
		conn, err := ln.Accept()
        if err != nil {
			log.Printf("Error: %s\n", err)
		}

		go handleRequest(conn)
    }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	fmt.Printf(".")
	conn.Write([]byte("OK."))
	conn.Close()
  }
