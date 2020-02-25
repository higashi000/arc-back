package slack

import (
	"log"

	"github.com/slack-go/slack"
)

func UserList(names []string) []slack.User {
	api := Api()

	users, err := api.GetUsers()
	if err != nil {
		log.Fatal(err)
	}

	var target []slack.User

	for _, e := range users {
		for _, targetName := range names {
			if e.RealName == targetName {
				target = append(target, e)
				break
			}
		}
	}

	return target
}
