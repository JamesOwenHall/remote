package remote

import (
	"encoding/gob"
	"fmt"
	"net"
)

// Remote is a server that can listen for commands.  To use it, create an
// instance with New.
type Remote struct {
	Commands map[string]func([]string) string
	listener net.Listener
}

// New returns an initialized Remote, or a non-nil error if there was a
// problem setting up the socket.
func New(path string) (*Remote, error) {
	result := Remote{
		Commands: make(map[string]func([]string) string),
	}

	var err error
	result.listener, err = net.Listen("unix", path)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ListenAndServe starts the server.  This function blocks until a call to StopListening is made, so it is often useful to start it in a new goroutine.  Always call StopListening() when you're done listening for events (such as when quitting the application) so that the socket gets deleted.
func (r *Remote) ListenAndServe() {
	// Loop until we can't accept new connections
	for conn, err := r.listener.Accept(); err == nil; conn, err = r.listener.Accept() {
		go r.serve(conn)
	}

	r.listener.Close()
}

// StopListening stops the server.  This should always be called once the server is no longer needed so that the socket gets deleted.
func (r *Remote) StopListening() {
	r.listener.Close()
}

func (r *Remote) serve(conn net.Conn) {
	defer conn.Close()

	message, err := readCommand(conn)
	if err != nil {
		return
	}

	if len(message) == 0 {
		return
	}

	if handler, ok := r.Commands[message[0]]; ok {
		fmt.Fprintln(conn, handler(message[1:]))
	} else {
		fmt.Fprintln(conn, message[0], "command does not exist")
	}
}

func readCommand(conn net.Conn) ([]string, error) {
	decoder := gob.NewDecoder(conn)
	var result []string
	err := decoder.Decode(&result)
	return result, err
}
