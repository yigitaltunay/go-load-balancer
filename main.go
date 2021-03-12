package main

import (
	ybalancer "github.com/yigitaltunay/go-load-balancer/ybalancer"
)

var balancer *ybalancer.LoadObject
func init() {
	listenAddress := "localhost:4200" // listen address
	server        := []string{ // server list
		"localhost:8080",
	}
	balancer = ybalancer.Create(listenAddress, server)
}

func main() {
	balancer.BalancerStart() // TCP load balancer
	balancer.ReverseProxyStart()  // Reverse Proxy 
}
