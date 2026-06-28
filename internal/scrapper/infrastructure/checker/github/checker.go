package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/domain"
)

type Checker struct {
	client *http.Client
	token  string
}

func New(client *http.Client, token string) *Checker {
	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}

	return &Checker{
		client: client,
		token:  token,
	}
}

func (c *Checker) Supports(rawURL string) bool {
	u, err := url.Parse(rawURL)
	if err != nil {
		return false
	}

	host := strings.ToLower(u.Host)
	return host == "github.com" || host == "www.github.com"
}

func (c *Checker) Check(ctx context.Context, link domain.TrackedLink) (application.CheckResult, error) {
	owner, repo, err := parseRepoURL(link.URL)
	if err != nil {
		return application.CheckResult{}, err
	}

	apiURL := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/commits?per_page=1",
		owner,
		repo,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return application.CheckResult{}, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return application.CheckResult{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return application.CheckResult{}, fmt.Errorf("github api returned status %d", resp.StatusCode)
	}

	var commits []struct {
		SHA    string `json:"sha"`
		Commit struct {
			Message   string `json:"message"`
			Committer struct {
				Date time.Time `json:"date"`
			} `json:"committer"`
		} `json:"commit"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&commits); err != nil {
		return application.CheckResult{}, err
	}

	if len(commits) == 0 {
		return application.CheckResult{
			HasUpdates: false,
		}, nil
	}

	lastCommit := commits[0]
	lastCommitTime := lastCommit.Commit.Committer.Date

	if !lastCommitTime.After(link.LastUpdatedAt) {
		return application.CheckResult{
			HasUpdates: false,
		}, nil
	}

	return application.CheckResult{
		HasUpdates:   true,
		NewUpdatedAt: lastCommitTime,
		Description:  fmt.Sprintf("new commit in %s/%s: %s", owner, repo, lastCommit.Commit.Message),
	}, nil
}

func parseRepoURL(rawURL string) (string, string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", "", err
	}

	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(parts) < 2 {
		return "", "", fmt.Errorf("invalid github repo url: %s", rawURL)
	}

	owner := parts[0]
	repo := parts[1]
	repo = strings.TrimSuffix(repo, ".git")

	return owner, repo, nil
}
