package main

import (
	"bytes"
	"fmt"
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

func commit(message string, push bool) {
	repoCheck := getCurDir()
	if repoCheck == "" {
		return
	}

	trackedFiles := getModifiedFiles()
	unTrackedFiles := getUntrackedFiles()
	for _, tf := range trackedFiles {
		if tf == "" {
			continue
		}

		commitFile(tf, false)
	}

	for _, uf := range unTrackedFiles {
		if uf == "" {
			continue
		}

		commitFile(uf, true)
	}

	cmd := exec.Command("git", "commit", "-m", message)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	if push {
		_, hasUpstream := getUpstreamUrl()
		if !hasUpstream {
			return
		}

		cmd = exec.Command("git", "push")
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}

func getUntrackedFiles() []string {
	var out bytes.Buffer
	cmd := exec.Command("git", "ls-files", "--others", "--exclude-standard")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	outStr := out.String()
	if outStr == "" {
		return make([]string, 0)
	}

	files := strings.Split(outStr, "\n")
	return files
}

func getModifiedFiles() []string {
	var out bytes.Buffer
	cmd := exec.Command("git", "ls-files", "--modified")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	outStr := out.String()
	if outStr == "" {
		return make([]string, 0)
	}

	files := strings.Split(outStr, "\n")
	return files
}

func commitFile(file string, isUntracked bool) {
	var choice string
	if isUntracked {
		fmt.Printf("Commit untracked file: '%s'? (y/n) ", file)
		fmt.Scan(&choice)
	}

	if !isUntracked || choice == "y" || choice == "Y" {
		cmd := exec.Command("git", "add", file)
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}
