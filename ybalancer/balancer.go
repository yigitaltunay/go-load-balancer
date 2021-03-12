package ybalancer

var counter int // not working

// This part is what is required for preparation..
func Create(listenAddress string, server []string) *LoadObject {
	p := LoadObject{Addrs: listenAddress, Serv: server}
	return &p
}

func ChooseRandom(serverelement []string) string {
	host := serverelement[counter%len(serverelement)]
	return host
}

func (e *LoadObject) GetListenAddress() string {
	return e.Addrs
}
