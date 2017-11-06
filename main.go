package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"flag"
	"bytes"
)

func main() {
	fmt.Println("ftp-client")

	username := flag.String("username", "", "remote host username")
	password := flag.String("password", "", "remote host password")
	host := flag.String("host", "", "remote host")
	flag.Parse()

	sshConfig := &ssh.ClientConfig{
		User: *username,
		Auth: []ssh.AuthMethod{
			ssh.Password(*password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	connection, err := ssh.Dial("tcp", *host + ":22", sshConfig)
	if err != nil {
		fmt.Printf("Failed to dial: %s", err)
		return
	}

	session, err := connection.NewSession()
	if err != nil {
		fmt.Printf("Failed to create session: %s", err)
		return
	}

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())

	defer session.Close()
}