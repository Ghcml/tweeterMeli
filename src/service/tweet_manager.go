package service

import (
	"fmt"

	"github.com/tweeterMeli/src/domain"
)

var Tweet *domain.Tweet

func PublishTweet(tweeet *domain.Tweet) error {

	if tweeet.User == "" {
		return fmt.Errorf("user is required")
	}
	if tweeet.Text == "" {
		return fmt.Errorf("text is required")
	}
	if len(tweeet.Text) > 140 {
		return fmt.Errorf("text exceeds 140 characters")
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
