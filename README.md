<div align="center">
<p>
    <img width="80" src="./logo.png">
</p>

<h1>看门狗 - 您的服务状态管家</h1>

[文档](docs.md) |
[变更日志](CHANGELOG.md) |
[贡献代码](CONTRIBUTING.md)

[![Build Status](https://github.com/axetroy/watchdog/workflows/ci/badge.svg)](https://github.com/axetroy/watchdog/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/watchdog)](https://goreportcard.com/report/github.com/axetroy/watchdog)
![Latest Version](https://img.shields.io/github/v/release/axetroy/watchdog.svg)
![License](https://img.shields.io/github/license/axetroy/watchdog.svg)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/watchdog.svg)

</div>

看门狗，一个服务监听者，监听各种协议的服务是否在线，然后通过各种渠道通知开发者。

### 从源码构建

Make sure you have `Golang@v1.16.x` installed.

```shell
$ git clone https://github.com/axetroy/watchdog.git $GOPATH/src/github.com/axetroy/watchdog
$ cd $GOPATH/src/github.com/axetroy/watchdog
$ make build
```

### 测试

```bash
$ make test
```

### 开源许可

The [MIT License](LICENSE)
