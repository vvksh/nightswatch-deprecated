package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

const webhook = "https://hooks.slack.com/services/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"

type SlackRequestBody struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func SendSlackNotification(msg string) error {

	slackBody, _ := json.Marshal(SlackRequestBody{Text: msg, Type: "mrkdwn"})
	req, err := http.NewRequest(http.MethodPost, webhook, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}
