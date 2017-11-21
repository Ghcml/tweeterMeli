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
func TestCleanTestIsClean(t *testing.T) {
	tweet := "This is my first tweet" //Creo un Tweet
	service.PublishTweet(tweet)       //Lo publico
	service.CleanTweet()              //Despues limpio el ulitmo Tweet
	if service.GetTweet() == tweet {
		t.Error("No funciona el tweet")
	}
}
