package ybalancer

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/fatih/color"
)

func (e *LoadObject) BalancerStart() {
	notice := color.New(color.Bold, color.FgGreen).PrintlnFunc()
	notice("[INFO] ListenAdress " + e.Addrs + " started ")
	listener, err := net.Listen("tcp", e.Addrs)
	if err != nil {
		notice("[ERROR] " + err.Error() + " ")
		return
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

// Rule area
func (e *LoadObject) AddRule() {}

func Rule(c net.Conn) { handleRequest(c) }

func handleRequest(conn net.Conn) {
	// Buffer that holds incoming information
	buf := make([]byte, 1024)
	len, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	s := string(buf[:len])
	fmt.Println("Stuff", s)

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
