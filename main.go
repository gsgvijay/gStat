package main

import (
	"flag"
	"fmt"
	"github.com/gookit/color"
	"strings"
)

func main() {
	userPtr := flag.String("user", "", "Git user name")
	repoPtr := flag.String("repo", "", "Repository")
	prNumPtr := flag.Int("pr", -1, "Pull Request number")
	flag.Parse()

	user, repos := nullCheck(userPtr, repoPtr)

	for _, repo := range(repos) {
		var nums []int
		if *prNumPtr != -1 {
			nums = make([]int, 1)
			nums[0] = *prNumPtr
		} else {
			nums = getPRNumsForRepo(user, repo)
		}

		for _, prNum := range nums {
			fmt.Printf("Comments for User: %s; Repo: %s; PR Num: %d\n", user, repo, prNum)
			comments := getPRComments(user, repo, prNum)
			if len(comments) == 0 {
				return
			}

			files := getFilesInPR(user, repo, prNum)
			for _, comment := range comments {
				contents := getFileContents(files[comment.File])
				context := getContext(contents, comment.Line, comment.Comment)
				fmt.Println(context)
				fmt.Println()
			}

			fmt.Println()
		}
	}
}

func getContext(fileContents string, lineNumber int, comment string) string {
	contents := strings.Split(fileContents, "\n")
	green := color.FgGreen.Render
	blue := color.FgBlue.Render

	context := ""
	start := lineNumber - 3
	if start < 0 {
		start = 0
	}

	end := lineNumber + 1
	if end >= len(contents) {
		end = len(contents) - 1
	}

	for i:=start; i<=end; i++ {
		if i == lineNumber-1 {
			context = context + blue(contents[i]) + "\t\t" + green("<--" + comment) + "\n"
		} else {
			context = context + contents[i] + "\n"
		}
	}

	return context
}

func nullCheck(userPtr *string, repoPtr *string) (string, []string) {
	user := *userPtr
	repo := *repoPtr
	if user == "" {
		user = getGitUser()
	}

	if repo == "" {
		repo = getCurDir()
	}

	if user == "" {
		fmt.Print("Username couldn't be determined. User: ")
		fmt.Scan(&user)
	}

	user = strings.TrimRight(user, "\r\n")

	var repos []string
	if repo == "" {
		allRepos := getUserRepos(user)
		repos = make([]string, len(allRepos))
		for i, ur := range(allRepos) {
			repos[i] = ur.RepoName
		}
	} else {
		repos = make([]string, 1)
		repos[0] = repo
	}

	return user, repos
}

func getPRNumsForRepo(user string, repo string) []int {
	prNums := getPRNums(user, repo)
	nums := make([]int, len(prNums))
	for i, n := range prNums {
		nums[i] = n.Number
	}

	return nums
}
