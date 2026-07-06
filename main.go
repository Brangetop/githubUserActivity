package main

import (
	"context"
	"fmt"
	"os"

	"brange.net/githubuseractivity/api"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "github-activity",
		Usage: "Показывает недавнюю активность пользователя на GitHub",

		Action: func(ctx context.Context, cmd *cli.Command) error {
			username := cmd.Args().Get(0)

			if username == "" {
				return fmt.Errorf("ошибка: необходимо указать имя пользователя. Пример: github-activity Brangetop")
			}

			err := api.FetchGitHubUserActivity(username)
			if err != nil {
				return err
			}

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
