package service

var tweet string

func PublishTweet(tweeet string) {
	tweet = tweeet
}

func GetTweet() string {
	return tweet
}
