[![Build Status](https://github.com/axetroy/watchdog/workflows/ci/badge.svg)](https://github.com/axetroy/watchdog/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/watchdog)](https://goreportcard.com/report/github.com/axetroy/watchdog)
![Latest Version](https://img.shields.io/github/v/release/axetroy/watchdog.svg)
![License](https://img.shields.io/github/license/axetroy/watchdog.svg)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/watchdog.svg)

## watchdog

> watchdog

### Usage

```bash

```

### Installation

If you have installed nodejs, you can install it via npm

```bash
npm install @axetroy/watchdog -g
```

If you are using Linux/macOS. you can install it with the following command:

```shell
# install latest version
curl -fsSL https://raw.githubusercontent.com/axetroy/watchdog/master/install.sh | bash
# or install specified version
curl -fsSL https://raw.githubusercontent.com/axetroy/watchdog/master/install.sh | bash -s v1.3.0
# or install from gobinaries.com
curl -sf https://gobinaries.com/axetroy/watchdog@v1.3.0 | sh
```

Or

Download the executable file for your platform at [release page](https://github.com/axetroy/watchdog/releases)

### Build from source code

Make sure you have `Golang@v1.16.x` installed.

```shell
$ git clone https://github.com/axetroy/watchdog.git $GOPATH/src/github.com/axetroy/watchdog
$ cd $GOPATH/src/github.com/axetroy/watchdog
$ make build
```

### Test

```bash
$ make test
```

### License

The [MIT License](LICENSE)
