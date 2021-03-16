package main

import (
	"github.com/axetroy/watchdog/proto"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

func main() {
	var (
		err error
	)

	if err = proto.PingHTTP("http://baidu.com"); err != nil {
		panic(err)
	}
}
