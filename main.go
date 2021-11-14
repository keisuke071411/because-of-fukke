package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func main() {
	envErr := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))

	if envErr != nil {
		fmt.Println(envErr)
	}

	tkn := os.Getenv("SLACK_BOT_TOKEN")
	api := slack.New(tkn)

	http.HandleFunc("/slack/events", func(w http.ResponseWriter, r *http.Request) {
		// リクエスト内容を取得
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		// イベント内容を取得
		eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		// イベント内容によって処理を分岐
		switch eventsAPIEvent.Type {
			case slackevents.URLVerification: // URL検証の場合の処理
				var res *slackevents.ChallengeResponse

				if err := json.Unmarshal(body, &res); err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)

					return
				}
				w.Header().Set("Content-Type", "text/plain")
				if _, err := w.Write([]byte(res.Challenge)); err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)

					return
				}

			case slackevents.CallbackEvent: // コールバックイベントの場合の処理
				innerEvent := eventsAPIEvent.InnerEvent
				fmt.Println("innerEvent", innerEvent.Type)

				// イベントタイプで分岐
				switch event := innerEvent.Data.(type) {
					case *slackevents.MessageEvent: 
						ref := slack.NewRefToMessage(event.Channel, event.TimeStamp)
						reg_name := regexp.MustCompile(`ふっけ`)
						reg_factor := regexp.MustCompile(`のせい`)
						reg_special := regexp.MustCompile(`ありがとう`)

						if txt := reg_name.FindString(event.Text); txt != "" {
							api.AddReaction("fukke0906",ref)
							return
						}

						if txt := reg_factor.FindString(event.Text); txt != "" {
							api.AddReaction("fukke0906",ref)
							return
						}

						if txt := reg_special.FindString(event.Text); txt != "" {
							api.AddReaction("lovefukke0906",ref)
							return
						}
				}
		}
	})

	log.Println("[INFO] Server listening")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Fatal(err)
	}
}
