package protocol

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

func PingHTTP(ctx context.Context, addr string) error {
	req, err := http.NewRequest(http.MethodHead, addr, nil)

	if err != nil {
		return errors.WithStack(err)
	}

	req = req.WithContext(ctx)

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return errors.WithStack(err)
	}

	defer res.Body.Close()

	if res.StatusCode >= http.StatusOK {
		return nil
	}

	return errors.New(res.Status)
}
