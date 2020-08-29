package gitlaborm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGitlabRepoURL(t *testing.T) {
	testCases := []struct {
		gitlabRepoURL       string
		expectedPid         string
		expectedBaseURL     string
		expectedErrorString string
	}{
		{
			"https://git.example.com/group/project",
			"group/project",
			"https://git.example.com/api/v4",
			"",
		},
		{
			"not a url",
			"",
			"",
			"Invalid gitlab url does not contain host: not a url",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.gitlabRepoURL, func(t *testing.T) {
			actualPid, actualBaseURL, actualError := parseGitlabRepoURL(tc.gitlabRepoURL)
			actualErrorString := ""
			if actualError != nil {
				actualErrorString = actualError.Error()
			}
			assert := assert.New(t)
			assert.Equal(tc.expectedPid, actualPid)
			assert.Equal(tc.expectedBaseURL, actualBaseURL)
			assert.Equal(tc.expectedErrorString, actualErrorString)
		})
	}

}
