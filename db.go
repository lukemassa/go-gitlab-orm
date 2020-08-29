package gitlaborm

import (
	"fmt"

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
	pid    int
}

// Record a record in the database
type Record interface {
	ID() string
}

// Connect to DB
func Connect(client *gitlab.Client, pid int, config Config) (*DB, error) {

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

func (db *DB) Ping() string {
	commits, _, err := db.client.Commits.ListCommits(db.pid, &gitlab.ListCommitsOptions{})
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return commits[0].Message
}
