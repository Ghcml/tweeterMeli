package service_test

import (
	"testing"

	"github.com/tweeterMeli/src/domain"
	"github.com/tweeterMeli/src/service"
)

func TestCleanTestIsClean(t *testing.T) {
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet := domain.NewTweet(user, text)
	service.PublishTweet(tweet)
	service.CleanTweet() //Despues limpio el ulitmo Tweet

	if service.GetTweet() == tweet {
		t.Error("No funciona el tweet")
	}
}
func TestPublishedTweetIsSaved(t *testing.T) {
	//Initializition
	var tweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)
	println(tweet)
	//Operation
	service.PublishTweet(tweet)

	//Validation
	publishedTweet := service.GetTweet()

	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, publishedTweet.User, publishedTweet.Text)
	}
	if publishedTweet.Date == nil {
		t.Error("Expected date can`t be nill")
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	//Initialization
	var tweet *domain.Tweet

	var user string = "hola"

	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	//operation
	var err error
	err = service.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	//Initialization
	var tweet *domain.Tweet

	user := "grupoesfera"
	var text string

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	err = service.PublishTweet(tweet)

	//Validation
	if err == nil {
		t.Error("Expected Error")
		return
	}
	if err.Error() != "text is required" {
		t.Error("Expected error is text is requiered")
	}
}
