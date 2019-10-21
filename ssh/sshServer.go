package main

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"github.com/gliderlabs/ssh"
)

func main() {
	ssh.Handle(func(s ssh.Session) {

		io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))
		io.WriteString(s, fmt.Sprintf("Hello 2%s\n", s.User()))
		io.WriteString(s, fmt.Sprintf("Hello 3%s\n", s.User()))

		text, err := bufio.NewReader(s).ReadString('\n')
		if err != nil {
			panic("GetLines: " + err.Error())
		}

		io.WriteString(s, fmt.Sprintf("ton texte %s\n", text))
	})

	log.Println("starting ssh server on port 2223...")
	log.Fatal(ssh.ListenAndServe(":2223", nil))
}
