package main

type InitDetails struct {
	GitUser string
	GitRepo string
}

type RepoInfo struct {
	Id	int `json:"id"`
	NodeId string `json:"node_id"`
	Name string `json:"name"`
	IsPrivate bool `json:"private"`
	Owner OwnerInfo `json:"owner"`
}

type OwnerInfo struct {
	Login string `json:"login"`
	Id int `json:"id"`
	Url string `json:"url"`
}

type IssueInfo struct {
	Url string `json:"url"`
	HtmlUrl string `json:"html_url"`
	IssueUrl string `json:"issue_url"`
	Id int `json:"id"`
	Created string `json:"created_at"`
	Updated string `json:"updated_at"`
	AuthorAssociation string `json:"author_association"`
	Body string `json:"body"`
}

type PRFiles struct {
	FileName string `json:"filename"`
	ContentUrl string `json:"contents_url"`
}

type PRFileContents struct {
	FileName string `json:"name"`
	FilePath string `json:"path"`
	Content string `json:"content"`
	Encoding string `json:"encoding"`
}

type PRComments struct {
	File string `json:"path"`
	Line int `json:"position"`
	User PRUser `json:"user"`
	Comment string `json:"body"`
	Created string `json:"created_at"`
	Updated string `json:"updated_at"`
}

type PRUser struct {
	UserName string `json:"login"`
	AvatarUrl string `json:"avatar_url"`
}

type PullRequest struct {
	Url string `json:"url"`
}

type PRNum struct {
	Number int `json:"number"`
}

type UserRepo struct {
	RepoName string `json:"name"`
}
