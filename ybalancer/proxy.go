package ybalancer

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func (e *LoadObject) ReverseProxyStart() {
	backendFind := ChooseRandom(e.Serv)
	remote, err := url.Parse(backendFind)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))
	err = http.ListenAndServe(e.Addrs, nil)
	if err != nil {
		panic(err)
	}
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p.ServeHTTP(w, r)
	}
}
