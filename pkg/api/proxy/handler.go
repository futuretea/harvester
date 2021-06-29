package proxy

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httputil"
)

// Handler proxies requests to the rancher service
type Handler struct {
	Scheme string
	Host string
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	director := func(r *http.Request) {
		r.URL.Scheme = "https"
		r.URL.Host = "rancher.cattle-system"
		r.URL.Path = req.URL.Path
		if h.Host != "" {
			r.URL.Host = h.Host
		}
		if h.Scheme != "" {
			r.URL.Scheme = h.Scheme
		}
		r.Host = r.URL.Host
		r.Header.Set("Origin", fmt.Sprintf("%s://%s", "https", r.Host))
	}
	httpProxy := &httputil.ReverseProxy{
		Director: director,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	httpProxy.ServeHTTP(rw, req)
}
