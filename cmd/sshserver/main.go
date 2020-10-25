package main

import (
	"io"
	"log"

	"github.com/gliderlabs/ssh"
)

func main() {
	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, "Hello world from Awan\n")
	})

	log.Fatal(ssh.ListenAndServe(":2220", nil))
}
