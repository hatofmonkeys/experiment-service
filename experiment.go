package main

import (
	"math/rand"
	"net/url"
	"strings"
)

func Specimen(srcUrl *url.URL, ratio int) (*url.URL, error) {
	modifiedUrl, err := url.Parse(srcUrl.String())
	hostComponents := strings.Split(modifiedUrl.Host, ".")

	magicRandomNumber := rand.Intn(102)

	if magicRandomNumber > ratio {
		modifiedUrl.Host = hostComponents[0] + "-b." + strings.Join(hostComponents[1:], ".")
	} else {
		modifiedUrl.Host = hostComponents[0] + "-a." + strings.Join(hostComponents[1:], ".")
	}
	return modifiedUrl, err
}
