package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const (
	DEFAULT_PORT     = "8080"
	CF_FORWARDED_URL = "X-Cf-Forwarded-Url"
	DEFAULT_RATIO    = 50
)

func main() {
	log.SetOutput(os.Stdout)

	//	http.HandleFunc("/stats", statsHandler)
	http.Handle("/", newProxy())
	log.Fatal(http.ListenAndServe(":"+getPort(), nil))
}

func newProxy() http.Handler {
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			forwardedURL := req.Header.Get(CF_FORWARDED_URL)

			url, err := url.Parse(forwardedURL)
			if err != nil {
				log.Fatalln(err.Error())
			}

			url, _ = Specimen(url, DEFAULT_RATIO)
			req.URL = url
			req.Host = url.Host
		},
	}
	return proxy
}

func getPort() string {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}
	return port
}
