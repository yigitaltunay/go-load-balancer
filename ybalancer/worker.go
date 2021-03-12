package ybalancer

import (
	"io"
	"log"
	"net"

	"github.com/fatih/color"
)

func (e *LoadObject) Start() {
	notice := color.New(color.Bold, color.FgGreen).PrintlnFunc()
	notice("[INFO] ListenAdress " + e.Addrs + " started")
	listener, err := net.Listen("tcp", e.Addrs)
	if err != nil {
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		
		if err != nil {
			log.Println(err)
		}

		backendFind := ChooseRandom(e.Serv)
		go Proxy(backendFind, conn)
	}

}

// Changes like control and udp will be made here..
func Proxy(backend string, c net.Conn) {
	backendconnection, err := net.Dial("tcp", backend)
	if err != nil {
		log.Println(err)
	}

	// thread
	go io.Copy(backendconnection, c)
	go io.Copy(c, backendconnection)

}
