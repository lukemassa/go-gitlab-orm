package gitlaborm

import "github.com/xanzy/go-gitlab"

// Config for the orm
type Config struct {
	User int
}

// DB object to make database calls
type DB struct {
	config Config
	client *gitlab.Client
}

// New DB
func New(client *gitlab.Client, config Config) (*DB, error) {

	ret := DB{
		config: config,
		client: client,
	}

	return &ret, nil
}
