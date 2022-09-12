package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	// range over event & print all the values
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
	}
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	/* Development
	// Set environment variables
	 os.Setenv("SLACK_BOT_TOKEN", "") //oauth token
	 os.Setenv("SLACK_APP_TOKEN", "") //socket token

	 botToken := goDotEnvVariable("SLACK_BOT_TOKEN")
	 fmt.Println("SLACK BOT TOKEN", botToken)
	 fmt.Printf("%s uses %s\n", os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	*/

	//init bot
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	// slack bot event of message ping, response pong
	bot.Command("ping", &slacker.CommandDefinition{
		Handler: func(botXtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
