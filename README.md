# Ybalancer
##### Golang Basic TCP Load Balancer
[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://github.com/yigitaltunay/go-load-balancer)

**Ybalancer**,  explains the load balancer  with Golang at a basic level. In some projects, the proxy can meet the connection request. Transmits incoming requests to the appropriate server.

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

```

![alt text](https://i.ibb.co/mJqDTJP/image.png?raw=true "Ybalancer")

### Todos

 - flag integration will be made 
 - server availability info

License
----
MIT



