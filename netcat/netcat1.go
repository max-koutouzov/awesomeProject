package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	address := flag.String("address", "", "-address localhost:8000")
	flag.Parse()
	conn, err := net.Dial("tcp", *address)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
