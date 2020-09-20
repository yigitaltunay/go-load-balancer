package main

import (
	ybalancer "go-load-balancer/ybalancer"
	"log"
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
