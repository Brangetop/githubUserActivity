package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func FetchGitHubUserActivity(user string) error {
	url := "https://api.github.com/users/" + user + "/events"

	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("error while getting http response", err)
		return errors.New("Error while gerring http response!" + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return errors.New("user not found" + user)
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("error, StatusCode:" + resp.Status)
	}

	var events []GitHubEvent
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return errors.New("Error while decoding JSON repsonse" + err.Error())
	}

	if len(events) == 0 {
		fmt.Println("user has no recent activity!")
		return nil
	}

	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			fmt.Println("Made ", len(event.Payload.Commits), "commits to repository ", event.Repo.Name)
		case "WatchEvent":
			fmt.Println("Starred repository ", event.Repo.Name)
		case "CreateEvent":
			fmt.Println("Created repository/branch ", event.Repo.Name)
		case "IssuesEvent":
			fmt.Println("Opened/closed issue in ", event.Repo.Name)
		default:
			fmt.Printf("Done %s in %s\n", event.Type, event.Repo.Name)
		}
	}

	return nil
}
