package main

import (
	"log"
	"net/http"
	"strings"
	"sync"
)

// Proxy contains and servers the handlers for each hostname
type Proxy struct {
	hostMap sync.Map
}

// Host contains the settings for each host
type Host struct {
	Target        string
	SetCookiePath bool
}

// Handle adds a handler if it doesn't exist
func (proxy *Proxy) Handle(host string, handler *ProxyHandler) {
	proxy.hostMap.Store(host, handler)
}

// Exists returns whether there is an
func (proxy *Proxy) Exists(host, target string) bool {
	item, ok := proxy.hostMap.Load(host)
	if !ok {
		return false
	}
	return item.(*ProxyHandler).TargetName == target
}

// ServeHTTP finds the handler if one exists and then returns the result
func (proxy *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if args.HSTS {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
	}
	// Match to hostname
	result, ok := proxy.hostMap.Load(r.Host)
	if ok {
		// Found a handler so serve
		handler := result.(*ProxyHandler)
		handler.Handler.ServeHTTP(w, r)
		return
	}
	// Match against the path prefix
	url := strings.Split(r.RequestURI, "/")
	if len(url) > 1 {
		result, ok = proxy.hostMap.Load("/" + url[1])
		if ok {
			// Found a handler so serve
			handler := result.(*ProxyHandler)
			handler.Handler.ServeHTTP(w, r)
			return
		}
	}
	// Hostname doesn't match so try wildcard
	result, ok = proxy.hostMap.Load("any")
	if ok {
		// Found a wildcard handler
		handler := result.(*ProxyHandler)
		handler.Handler.ServeHTTP(w, r)
	} else {
		http.Error(w, "Not found", 404)
	}

}

type proxyTransport struct {
	SetCookiePath     bool
	CapturedTransport http.RoundTripper
}

func (t *proxyTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	// Use the real transport to execute the request
	response, err := transport.RoundTrip(r)
	if err != nil {
		transport.(*http.Transport).CloseIdleConnections()
		log.Print("Unable to get response from target server: " + err.Error())
		return nil, err
	}
	if response.StatusCode >= 500 {
		transport.(*http.Transport).CloseIdleConnections()
	}
	if t.SetCookiePath {
		for name, values := range response.Header {
			if strings.EqualFold(name, "SET-COOKIE") {
				// Remove the current SET-COOKIE headers
				response.Header.Del("SET-COOKIE")
				for _, value := range values {
					parts := strings.Split(value, ";")
					// Update the cookie with a root path
					newSetCookie := parts[0] + "; Path=/"
					if len(parts) > 2 {
						newSetCookie += ";" + parts[2]
					}
					response.Header.Add("SET-COOKIE", newSetCookie)
				}
			}
		}
	}
	return response, err
}
