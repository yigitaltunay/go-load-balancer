# Ybalancer
##### Golang Basic TCP Load Balancer
[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)


<<<<<<< HEAD
**Ybalancer**,  explains the load balancer business with Golang at a basic level. In some projects, the proxy can meet the connection request. Transmits incoming requests to the appropriate server.
=======
**Ybalancer**,  explains the load balancer business with Golang at a basic level. In some projects, the proxy can meet the connection request. Transmits incoming requests to the appropriate server.
>>>>>>> 57c2c03258452202c4151b63adc11be169ec292c

### Installation
The first need [Go](https://golang.org/) installed (version 1.15+ is required), then you can use the below Go command to install Ybalancer.
```sh
$ go get -u github.com/yigitaltunay/go-load-balancer
```
Import it in your code:
```sh
import "github.com/yigitaltunay/go-load-balancer
```

### Build

```sh
$ go build
$ ./go-load-balancer
```
### How does it work?

```sh
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
		"localhost:8084",
		"localhost:8085",
	}
)
func main() {
	log.Println("Started Load Balancer")
	balancer := ybalancer.Create(listenAddress, server)
	balancer.Start()
}
```

![alt text](https://i.ibb.co/mJqDTJP/image.png?raw=true "Ybalancer")

### Todos

 - flag integration will be made 
 - server availability info

License
----
MIT



