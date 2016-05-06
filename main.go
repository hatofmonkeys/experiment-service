package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"

	"gopkg.in/redis.v3"
)

const (
	DEFAULT_PORT     = "8080"
	CF_FORWARDED_URL = "X-Cf-Forwarded-Url"
	DEFAULT_RATIO    = 50
)

func main() {
	log.SetOutput(os.Stdout)

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

			url, _ = Specimen(url, ratio())
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

func ratio() int {
	client := redis.NewClient(&redis.Options{
		Addr:     "xxxxxxxxxxxxx",
		Password: "xxxxxxxxxxxxx", // no password set
		DB:       0,               // use default DB
	})

	strRatio, _ := client.Get("ratio").Result()
	intRatio, _ := strconv.Atoi(strRatio)

	return intRatio
}
