package main

import (
	ybalancer "github.com/yigitaltunay/go-load-balancer/ybalancer"
)



func main() {
	listenAddress := "localhost:4200" // listen address
	server        := []string{ // server list
		"localhost:8080",
	}
	balancer := ybalancer.Create(listenAddress, server)
	balancer.Start()
}
