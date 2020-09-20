# Ybalancer
##### Golang Basic TCP Load Balancer
[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)


**Ybalancer**,  xplains the load balancer business with Golang at a basic level. In some projects, the proxy can meet the connection request. Transmits incoming requests to the appropriate server.

### Installation

Ybalancer requires [Go](https://nodejs.org/) 1.15.2 version to run.
Install the dependencies and devDependencies and start the server.

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



