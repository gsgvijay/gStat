package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func getCommentsForIssue(user string, repo string, issue int) []IssueInfo {
	client := http.Client{}
	uri := "https://api.github.com/repos/" + user + "/" + repo + "/issues/" + strconv.Itoa(issue) + "/comments"
	req := getRequest(uri)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	} else if (resp.StatusCode != 200) {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var decoded []IssueInfo
	err = json.Unmarshal([]byte(body), &decoded)
	if err != nil {
		panic(err)
	}

	return decoded
}

func getFilesInPR(user string, repo string, prNum int) map[string]string {
	client := http.Client{}
	uri := "https://api.github.com/repos/" + user + "/" + repo + "/pulls/" + strconv.Itoa(prNum) + "/files"
	req := getRequest(uri)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	} else if resp.StatusCode != 200 {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var files []PRFiles
	err = json.Unmarshal([]byte(body), &files)
	if err != nil {
		panic(err)
	}

	fileMap := make(map[string]string)
	for _, file := range files {
		fileMap[file.FileName] = file.ContentUrl
	}

	return fileMap
}

func getFileContents(uri string) string {
	client := http.Client{}
	req := getRequest(uri)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	} else if resp.StatusCode != 200 {
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var fileContents PRFileContents
	err = json.Unmarshal([]byte(body), &fileContents)
	if err != nil {
		panic(err)
	}

	var contents string = ""
	if strings.EqualFold(fileContents.Encoding, "base64") {
		bContents, _ := b64.StdEncoding.DecodeString(fileContents.Content)
		contents = string(bContents)
	}

	return contents
}

func getPRComments(user string, repo string, prNum int) []PRComments {
	uri := "https://api.github.com/repos/" + user + "/" + repo + "/pulls/" + strconv.Itoa(prNum) + "/comments"
	return getPRCommentsFromUri(uri)
}

func getPRCommentsFromUri(uri string) []PRComments {
	client := http.Client{}
	req := getRequest(uri)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	} else if resp.StatusCode != 200 {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var comments []PRComments
	err = json.Unmarshal([]byte(body), &comments)
	if err != nil {
		panic(err)
	}

	return comments
}

func getPRs(user string, repo string) []PullRequest {
	uri := "https://api.github.com/repos/" + user + "/" + repo + "/pulls"
	client := http.Client{}
	req := getRequest(uri)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	} else if resp.StatusCode != 200 {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var pullRequests []PullRequest
	err = json.Unmarshal([]byte(body), &pullRequests)
	if err != nil {
		panic(err)
	}

	return pullRequests
}

func getUserRepos(user string) []UserRepo {
	uri := "https://api.github.com/users/" + user + "/repos"
	client := http.Client{}
	req := getRequest(uri)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	} else if resp.StatusCode != 200 {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var repos []UserRepo
	err = json.Unmarshal([]byte(body), &repos)
	if err != nil {
		panic(err)
	}

	return repos
}

func getPRNums(user string, repo string) []PRNum {
	uri := "https://api.github.com/repos/" + user + "/" + repo + "/pulls"
	client := http.Client{}
	req := getRequest(uri)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	} else if resp.StatusCode != 200 {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var prNums []PRNum
	err = json.Unmarshal([]byte(body), &prNums)
	if err != nil {
		panic(err)
	}

	return prNums
}
