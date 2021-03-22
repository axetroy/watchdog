package protocol

import (
	"errors"
	"net/http"
)

func PingHTTP(addr string) error {
	res, err := http.Head(addr)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode >= http.StatusOK {
		return nil
	}

	return errors.New(res.Status)
}
