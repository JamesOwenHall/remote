package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	var path string
	var help bool
	flag.StringVar(&path, "path", "", "the path to the unix socket")
	flag.StringVar(&path, "p", "", "same as -path")
	flag.BoolVar(&help, "help", false, "show help text")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}
	if path == "" {
		fmt.Fprintln(os.Stderr, `error: you must set the socket path, see "remotec --help"`)
		return
	}

	conn, err := net.Dial("unix", path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: can't connect to socket")
		return
	}
	defer conn.Close()

	err = sendCommand(conn, flag.Args())
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: can't send message")
		return
	}

	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: can't read response")
		return
	}
}

func sendCommand(conn net.Conn, message []string) error {
	encoder := gob.NewEncoder(conn)
	return encoder.Encode(message)
}
