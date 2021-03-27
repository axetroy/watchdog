package notify

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"github.com/axetroy/watchdog"
	"github.com/pkg/errors"
)

type WebhookResult struct {
	Url   string `json:"url"`
	Error error  `json:"error"`
}

func request(wg *sync.WaitGroup, ch chan WebhookResult, url string, body []byte) {
	var (
		err error
	)
	defer func() {
		ch <- WebhookResult{
			Url:   url,
			Error: err,
		}
	}()

	defer wg.Done()
	res, err := http.Post(url, "application/json", bytes.NewReader(body))

	if err != nil {
		err = errors.WithStack(err)
		return
	}

	defer res.Body.Close()
}

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

	messages := make(chan WebhookResult)
	results := make([]WebhookResult, 0)
	wg := sync.WaitGroup{}

	wg.Add(len(reporter.Target))

	for _, r := range reporter.Target {
		go request(&wg, messages, r, b)
	}

	for i := range messages {
		results = append(results, i)
		if len(results) == len(reporter.Target) {
			break
		}
	}

	errorList := make([]string, 0)

	for _, r := range results {
		if r.Error != nil {
			errorList = append(errorList, r.Error.Error())
		}
	}

	wg.Wait()

	if len(errorList) != 0 {
		return errors.New(strings.Join(errorList, ";"))
	} else {
		return nil
	}
}
