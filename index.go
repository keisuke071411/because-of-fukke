package main

import (
	"github.com/slack-go/slack"
)

func main() {
	const tkn string = "xoxp-534076003925-930421063989-2624919337831-7a0d646cbe5253d614779d05f40b1cea"
	c := slack.New(tkn)

		_, _, err := c.PostMessage("#three-sacred-treasures", slack.MsgOptionText("Hello World", true))
	if err != nil {
		panic(err)
	}
}

