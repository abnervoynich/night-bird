package utils

import (
	"github.com/abnervoynich/night-bird/app/utils/notifiers"
	"log"
)

func sendNotification(message string) error {

	return nil
}

func NotifyToSlack(message string) {
	sc := notifiers.SlackClient{
		WebHookUrl: "https://WEB_HOOK_URL",
		UserName:   "USER_NAME",
		Channel:    "CHANNEL_NAME",
	}
	sr := notifiers.SimpleSlackRequest{
		Text:      message,
		IconEmoji: ":ghost:",
	}
	err := sc.SendSlackNotification(sr)
	if err != nil {
		log.Fatal(err)
	}

	//To send a notification with status (slack attachments)
	sj := notifiers.SlackJobNotification{
		Text:      "This is attachment message",
		Details:   "details of the jobs",
		Color:     "warning",
		IconEmoji: ":hammer_and_wrench",
	}
	err = sc.SendJobNotification(sj)
	if err != nil {
		log.Fatal(err)
	}
}
