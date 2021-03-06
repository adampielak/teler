package alert

import (
	"github.com/slack-go/slack"
)

func toSlack(token string, channel string, color string, log map[string]string) {
	api := slack.New(token)
	reason := slack.Attachment{
		AuthorName: ":warning: teler Alert",
		Title:      log["category"],
		Color:      color,
	}
	request := slack.Attachment{
		Title: "Request",
		Text:  log[log["element"]],
		Color: color,
	}
	fields := slack.Attachment{
		Color: color,
		Fields: []slack.AttachmentField{
			{
				Title: "Date",
				Value: log["time_local"],
				Short: true,
			},
			{
				Title: "Referrer",
				Value: log["http_referer"],
				Short: true,
			},
			{
				Title: "User-Agent",
				Value: log["http_user_agent"],
				Short: true,
			},
			{
				Title: "Status code",
				Value: log["status"],
				Short: true,
			},
			{
				Title: "Bytes sent",
				Value: log["body_bytes_sent"],
				Short: true,
			},
		},
	}

	// TODO: Displays an error if it does not exceed the rate-limit
	// nolint:errcheck
	api.PostMessage(
		channel,
		slack.MsgOptionAttachments(reason, request, fields),
		slack.MsgOptionAsUser(true),
	)
}
