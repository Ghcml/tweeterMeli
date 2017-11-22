package main

import (
	"github.com/abiosoft/ishell"
	"github.com/tweeterMeli/src/domain"
	"github.com/tweeterMeli/src/service"
)

func main() {
	shell := ishell.New()
	shell.SetPrompt("Tweeter >>")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			c.Print("Type your username: ")
			user := c.ReadLine()

			c.Print("Write your tweet: ")
			text := c.ReadLine()

			tweet := domain.NewTweet(user, text)
			err := service.PublishTweet(tweet)
			if err == nil {
				c.Print("Tweet sent\n")
			} else {
				c.Print("Error publishing tweet:", err)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "show a tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			tweet := service.GetTweet()
			c.Println(tweet)
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "cleanTweet",
		Help: "clean last tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			service.CleanTweet()
			c.Println("Borraste el ultimo Tweet")
			return
		},
	})
	shell.Run()
}
