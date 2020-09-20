package loadbalancer

import (
	"io"
	"log"
	"net"
)

type Object struct {
	listenAddress string
	server        []string
}

var counter int

func Create(listenAddress string, server []string) *Object {
	p := Object{listenAddress: listenAddress, server: server}
	return &p
}

func (e *Object) Start() {

	listener, err := net.Listen("tcp", e.listenAddress)
	if err != nil {

	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		backendFind := ChooseRandom(e.server)
		go proxy(backendFind, conn)
	}

}

func proxy(backend string, c net.Conn) {
	backendconnection, err := net.Dial("tcp", backend)
	if err != nil {
		log.Println(err)
	}

	// thread
	go io.Copy(backendconnection, c)
	go io.Copy(c, backendconnection)

}

func ChooseRandom(serverelement []string) string {
	host := serverelement[counter%len(serverelement)]
	return host
}

func (e *Object) GetListenAddress() string {
	return e.listenAddress
}
