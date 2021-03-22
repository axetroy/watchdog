package notify

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/axetroy/watchdog"
)

func Webhook(content string, reporter watchdog.Reporter) (err error) {
	type Body struct {
		Content string `json:"content"`
	}

	body := Body{
		Content: content,
	}

	b, err := json.Marshal(body)

	if err != nil {
		return nil
	}

	for _, r := range reporter.Target {
		res, err := http.Post(r, "application/json", bytes.NewReader(b))

		if err != nil {
			log.Printf("%+v\n", err.Error())
			continue
		}

		defer res.Body.Close()
	}

	return nil
}
