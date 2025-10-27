package internal

import "time"

type User struct {
	URL       string    `json:"html_url"`
	Username  string    `json:"login"`
	Repos     int       `json:"public_repos"`
	CreatedAt time.Time `json:"created_at"`
	Social    string    `json:"twitter_username"`
}
