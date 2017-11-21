package service_test

import (
	"testing"

	"github.com/tweeterMeli/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	//var tweet string = "This is my first tweet" //Declaro una variable con un tipo
	tweet := "This is my first tweet"
	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected tweet is", tweet)
	}
}
