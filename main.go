package main

import (
	"fmt"

	"brange.net/githubuseractivity/api"
)

func main() {
	err := api.FetchGitHubUserActivity("Brangetop")
	if err != nil {
		fmt.Println("Error while fetching", err)
	}
}
