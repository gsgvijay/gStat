package main

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
)

func getRepos(username string) []RepoInfo {
	client := http.Client{}
	uri := "https://api.github.com/users/" + username + "/repos"
	req := getRequest(uri)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var decoded []RepoInfo
	err = json.Unmarshal([]byte(body), &decoded)
	if err != nil {
		panic(err)
	}

	return decoded
}
