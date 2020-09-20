package main

import (
	"log"

	ybalancer "github.com/yigitaltunay/go-load-balancer/ybalancer"
)

var (
	listenAddress = "localhost:4200"
	counter       int
	server        = []string{
		"localhost:8081",
		"localhost:8082",
		"localhost:8083",
	}
)

func main() {

	log.Println("Started Load Balancer")
	balancer := ybalancer.Create(listenAddress, server)
	balancer.Start()

}
