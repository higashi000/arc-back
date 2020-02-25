package tweet

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func Tweet(tweetText string, userID []string) error {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))

	text := ""

	for _, e := range userID {
		text += e + " "
	}

	text += tweetText

	_, err := api.PostTweet(text, nil)
	if err != nil {
		return err
	}

	return nil
}
