package github

type GithubClient struct {
	url string
}
type GithubService interface {
	GetGists(username string) ([]Gist, error)
	Serve(port int)
}

type Gist struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Public      bool      `json:"public"`
	Url         string    `json:"url"`
	ForksUrl    string    `json:"forks_url"`
	CommitsUrl  string    `json:"commits_url"`
	NodeID      string    `json:"node_id"`
	GitPullUrl  string    `json:"git_pull_url"`
	GitPushUrl  string    `json:"git_push_url"`
	HTMLUrl     string    `json:"html_url"`
	Files       GistFiles `json:"files"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
	Owner       *User     `json:"owner,omitempty"`
	Comments    int       `json:"comments"`
	CommentsUrl string    `json:"comments_url"`
	User        *User     `json:"user,omitempty"`
	Truncated   bool      `json:"truncated"`
}

type GistFiles map[string]GistFile

type GistFile struct {
	Filename string `json:"filename"`
	Type     string `json:"type"`
	Language string `json:"language"`
	RawUrl   string `json:"raw_url"`
	Size     int    `json:"size"`
}

type User struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}
