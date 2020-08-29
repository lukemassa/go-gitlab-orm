package gitlaborm

import (
	"fmt"
	"strings"

	"net/url"

	"github.com/xanzy/go-gitlab"
)

// Config for the orm
type Config struct {
	User int
}

// DB object to make database calls
type DB struct {
	config Config
	client *gitlab.Client
	pid    string
}

// Record a record in the database
type Record interface {
	ID() string
}

func parseGitlabRepoURL(gitlabRepoURL string) (string, string, error) {
	parsedURL, err := url.Parse(gitlabRepoURL)
	if err != nil {
		return "", "", err
	}
	if parsedURL.Host == "" {
		return "", "", fmt.Errorf("Invalid gitlab url does not contain host: %s", gitlabRepoURL)
	}
	// Trim leading slash
	pid := strings.TrimLeft(parsedURL.Path, "/")
	baseURL := url.URL{
		Scheme: parsedURL.Scheme,
		Host:   parsedURL.Host,
		Path:   "/api/v4",
	}
	return pid, baseURL.String(), nil
}

// Connect to DB
func Connect(gitlabRepoURL string, token string, config Config) (*DB, error) {

	pid, baseURL, err := parseGitlabRepoURL(gitlabRepoURL)
	if err != nil {
		return nil, err
	}
	client, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseURL))
	if err != nil {
		return nil, err
	}

	ret := DB{
		config: config,
		client: client,
		pid:    pid,
	}

	return &ret, nil
}

// List all records in the database
//func (db *DB) List() []Record {
//    db.client.Branches.ListBranches()
//}

// Ping test connection to the database
func (db *DB) Ping() string {
	return fmt.Sprintf("Hello %s", db.client.BaseURL().User.Username())
}
