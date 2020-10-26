package main

import (
	"io"
	"log"

	"github.com/awanio/awan/internal/runtime"
	"github.com/gliderlabs/ssh"
)

func main() {

	runtime.Setup()

	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, "Hello world from Awan\n")
	})

	log.Fatal(ssh.ListenAndServe(":2220", nil))
}
