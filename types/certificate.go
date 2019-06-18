package data

import (
	"errors"
	"file"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
)

type Certificate struct {
	authority, certificate, key string
}

// --certificates=vault01=ca:/etc/certs/service.ca.crt,cert:/etc/certs/client.crt,key:/etc/certs/client.pem
func ParseCertificate(value string) (cert *Certificate, err error) {
	urls := strings.Split(value, ",")
	var authority, certificate, key string

	for _, s := range urls {
		u, err := url.Parse(s)
		if err != nil {
			continue
		}
		file, err := os.Open(u.Path)
		if err != nil {
			continue
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			continue
		}

		switch u.Scheme {
		case "ca":
			authority = data
		case "cert":
			certificate = data
		case "key":
			key = data
		}
	}

	return &Certificate{authority, certificate, key}, nil
}
