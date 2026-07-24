package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

var (
	host = "ws://localhost"
)

type TestConfig struct {
	clientCount int
	wg          *sync.WaitGroup
}

func DialServer(wg *sync.WaitGroup) {

	dialer := &websocket.Dialer{
		NetDial: func(network, addr string) (net.Conn, error) {
			return net.Dial("tcp4", addr) // force IPv4
		},
	}
	conn, _, err := dialer.Dial(fmt.Sprintf("%s%s", host, WSPort), nil)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		conn.Close()
		wg.Done()
	}()

	fmt.Println("connected to the server", conn.LocalAddr().String())
	time.Sleep(1 * time.Second)

}
func TestConnecting(t *testing.T) {
	go createWSServer()
	time.Sleep(1 * time.Second)

	tc := TestConfig{
		clientCount: 3,
		wg:          new(sync.WaitGroup),
	}

	tc.wg.Add(tc.clientCount)

	for range tc.clientCount {
		go DialServer(tc.wg)
	}
	tc.wg.Wait()
	fmt.Println("Exiting testing")

}
