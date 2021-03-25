package protocol

import (
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/pkg/errors"
)

func PingFTP(addr string) error {
	c, err := ftp.Dial(addr, ftp.DialWithTimeout(10*time.Second))

	if err != nil {
		return errors.WithStack(err)
	}

	if err := c.Quit(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
