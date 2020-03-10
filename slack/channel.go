package slack

import "log"

func GetChannelID(channelName string) string {
	api := Api()

	channels, err := api.GetChannels(false)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range channels {
		if e.Name == channelName {
			return e.Name
		}
	}

	return ""
}
