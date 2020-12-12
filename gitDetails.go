package main

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"regexp"
)

func getCurDir() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if !dirExists(path + "/.git") {
		return ""
	}

	var ss []string
	if runtime.GOOS == "windows" {
		ss = strings.Split(path, "\\")
	} else {
		ss = strings.Split(path, "/")
	}

	dirName := ss[len(ss) - 1]
	return dirName
}

func getGitUser() string {
	var out bytes.Buffer
	cmd := exec.Command("git", "config", "user.name")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	return out.String()
}

func getUpstreamUrl() (string, bool) {
	var out bytes.Buffer
	cmd := exec.Command("git", "remote", "-v")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	str := out.String()
	if len(str) == 0 {
		return "", false
	}

	r, _ := regexp.Compile("http.*.git")
	return r.FindString(str), true
}

func dirExists(path string) bool {
	stat, err := os.Stat(path)
	if err == nil {
		return stat.IsDir()
	}

	return false
}
