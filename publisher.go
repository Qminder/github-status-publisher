package main

import (
	"context"
	"fmt"
	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v29/github"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {

	if len(os.Args) < 4 {
		fmt.Println("Usage: context state description [url]")
		os.Exit(1)
	}

	owner, repo := getOwnerAndRepo()
	ref := getEnv("BUILDKITE_COMMIT")

	stateContext := os.Args[1]
	state := os.Args[2]
	description := os.Args[3]

	status := &github.RepoStatus{State: &state, Context: &stateContext, Description: &description}

	if len(os.Args) == 5 {
		status.TargetURL = &os.Args[4]
	}

	client := createGithubClient()
	_, _, err := client.Repositories.CreateStatus(context.Background(), owner, repo, ref, status)
	if err != nil {
		fmt.Println("Failed to publish status", err)
		os.Exit(1)
	}
}

func createGithubClient() *github.Client {
	appID, _ := strconv.ParseInt(getEnv("APP_ID"), 10, 64)
	installationID, _ := strconv.ParseInt(getEnv("INSTALLATION_ID"), 10, 64)
	privateKey := getEnv("GITHUB_APP_PRIVATE_KEY")

	itr, err := ghinstallation.New(http.DefaultTransport, appID, installationID, []byte(privateKey))
	if err != nil {
		fmt.Println("Failed to load the private key")
		os.Exit(1)
	}

	return github.NewClient(&http.Client{Transport: itr})
}

func getOwnerAndRepo() (string, string) {
	pattern := regexp.MustCompile(`([a-zA-Z]+)/([a-z]+)`)
	url := getEnv("BUILDKITE_REPO")
	matches := pattern.FindStringSubmatch(url)
	if len(matches) != 3 {
		fmt.Println("Value of BUILDKITE_REPO is unexpected")
		os.Exit(1)
	}

	return matches[1], matches[2]
}

func getEnv(name string) string {
	value, found := os.LookupEnv(name)
	if found == false {
		fmt.Printf("Please make sure %s environment variable is set\n", name)
		os.Exit(1)
	}
	return value
}
