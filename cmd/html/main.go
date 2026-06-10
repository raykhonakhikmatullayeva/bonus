package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		panic(err)
	}
	defer func(ln net.Listener) {
		err := ln.Close()
		if err != nil {
		}
	}(ln)
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}
func handle(conn net.Conn) {
	body := "Now: " + time.Now().Format("02.01.2006 15:04:05")
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return
	}
	reqLine := strings.SplitN(string(buf[:n]), "\r\n", 2)[0]
	parts := strings.Fields(reqLine)
	method, path := parts[0], parts[1]

	fmt.Printf("%s %s\n", method, path)
	writeResponse(conn, "200 ok", body)
}
func writeResponse(conn net.Conn, status, body string) {
	resp := fmt.Sprintf(
		`HTTP/1.1 %s
          Content-Type: text/html; charset=utf-8
          Content-Length: %d
          Connection: close

          %s`,
		status, len([]byte(body)), body,
	)

	conn.Write([]byte(resp))
}
