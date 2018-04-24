// Copyright 2018 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be found
// in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	var proxyStr string
	var responderStr string

	flag.StringVar(&proxyStr, "proxy_url", "", "Proxy URL")
	flag.StringVar(&responderStr, "responder_url", "", "Destination OCSP Responder URL")
	flag.Parse()

	if proxyStr == "" {
		log.Fatal("no Proxy URL set")
	}
	proxy, err := url.Parse(proxyStr)
	if err != nil {
		log.Fatal(err)
	}

	if responderStr == "" {
		log.Fatal("no OCSP Responder URL set")
	}
	responder, err := url.Parse(responderStr)
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe("127.0.0.1:8234", &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = responder.Scheme
			req.URL.Host = responder.Host
			req.URL.Path = responder.Path + req.URL.Path
			req.Host = responder.Host
		},
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	})
}
