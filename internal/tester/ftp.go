package tester

import (
	"log"
	"time"

	"goftp.io/server/v2"
)

func CreateFTPServer(addr string, cb func()) error {
	server, err := server.NewServer(&server.Options{
		Port: 8888,
		Perm: server.NewSimplePerm("test", "test"),
	})

	if err != nil {
		return err
	}

	defer func() {
		_ = server.Shutdown()
	}()

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	time.Sleep(time.Second * 3)

	cb()

	return nil
}
