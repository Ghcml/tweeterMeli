package service

import (
	"fmt"

	"github.com/tweeterMeli/src/domain"
)

var Tweet *domain.Tweet

func PublishTweet(tweeet *domain.Tweet) error {

	if tweeet.User == " " {
		return fmt.Errorf("user is required")
	}
	Tweet = tweeet
	return nil
}

func GetTweet() *domain.Tweet {
	return Tweet
}

func CleanTweet() {
	Tweet = nil
}
