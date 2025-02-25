package internal_scm

import (
	"github.com/jenkins-x/go-scm/scm/driver/github"
	"github.com/jenkins-x/go-scm/scm/transport"
	"net/http"
)

type GitHub struct {
	repo  string
	token string
}

// NewGitHub creates a new instance
func NewGitHub(repo, token string) *GitHub {
	return &GitHub{
		repo:  repo,
		token: token,
	}
}

func (r *GitHub) Release(version, commitish string, draft, prerelease bool) (err error) {
	client := github.NewDefault()
	client.Client = &http.Client{
		Transport: &transport.BearerToken{
			Token: r.token,
		},
	}
	err = release(client, r.repo, version, commitish, draft, prerelease)
	return
}
