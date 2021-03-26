package protocol

import (
	"context"
	"net"

	"github.com/pkg/errors"
)

func dialTCP(ctx context.Context, network, addr string) (net.Conn, error) {
	d := net.Dialer{}
	conn, err := d.DialContext(ctx, network, addr)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func PingTCP(ctx context.Context, addr string) error {
	_, err := net.ResolveTCPAddr("tcp", addr)

	if err != nil {
		return errors.WithStack(err)
	}

	conn, err := dialTCP(ctx, "tcp", addr)

	if err != nil {
		return errors.WithStack(err)
	}

	defer conn.Close()

	return nil
}
