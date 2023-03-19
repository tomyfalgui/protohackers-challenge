package echo

import (
	"fmt"
	"io"
	"log"
	"net"
)

func StartEcho() {
	addr := ":8080"
	ln, err := net.Listen("tcp", addr)
	defer ln.Close()
	if err != nil {
		log.Fatal("failed to start tcp server")
	}
	log.Println("Server is running on:", addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("failed to accept incoming tcp request")
			continue
		}
		go echo(conn)
	}
}

func echo(conn net.Conn) {
	defer conn.Close()
	fmt.Print(conn)
	io.Copy(conn, conn)
}
