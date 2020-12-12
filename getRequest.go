package main

import "net/http"

func getRequest(uri string) *http.Request {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")
	return req
}
