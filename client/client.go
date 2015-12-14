package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustcopy(conn, os.Stdout)
	mustcopy(os.Stdin, conn)
}

func mustcopy(r io.Reader, w io.Writer) {
	io.Copy(w, r)
}
