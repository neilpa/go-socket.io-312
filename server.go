package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

var (
	addr = flag.String("addr", "localhost:5432", "address listen/bind on")
)

func main() {
	flag.Parse()
	log.SetPrefix("server: ")
	log.SetFlags(log.LUTC | log.Lmicroseconds)

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("", func(s socketio.Conn) error {
		log.Println("connected:", s.ID())
		//s.Emit("ping", "connected: " + s.ID())
		return nil
	})
	server.OnEvent("", "foo", func(s socketio.Conn, msg string) {
		log.Println("recv'd foo, sending bar:", msg)
		s.Emit("bar", msg)
	})
	server.OnError("", func(e error) {
		log.Printf("error: %s\n", e)
	})
	server.OnDisconnect("", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})
	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	log.Println("Listening on", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
