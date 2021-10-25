package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	envErr := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
	
	if envErr != nil {
		fmt.Println(envErr)
	}

	tkn := os.Getenv("SLACK_BOT_TOKEN")
	c := slack.New(tkn)

	_, _, err := c.PostMessage("#three-sacred-treasures", slack.MsgOptionText("Hello World", true))
	if err != nil {
		panic(err)
	}
}

