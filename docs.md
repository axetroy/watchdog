# 看门狗文档

## 介绍

看门狗是一个 CLI 工具，用于监控各种服务的状态，它支持多种协议和多种通知渠道。

## 目录

<table>
    <tr><td width=25% valign=top>

- [介绍](#介绍)
- [安装](#安装)
- [快速开始](#快速开始)
- [FAQ](#faq)
  - [工作原理](#工作原理)
  - [路线图](#路线图)
  - [常见问题](#常见问题)

</td><td width=25% valign=top>

- [配置文件](#配置文件)
- [服务配置](#服务配置)
- [支持的协议](#支持的协议)
- [支持的通知渠道](#支持的通知渠道)
  - [通知协议](#通知协议)

</td><td width=25% valign=top>

- [使用例子](#使用例子)
  - [单服务](#单服务)
  - [多服务](#多服务)
  - [通知渠道](#通知渠道)
    - [微信通知](#微信通知)
    - [Webhook](#Webhook通知)
    - [邮件通知](#邮件通知)

</td><td width=25% valign=top>

- [开发](#开发)
  - [开发环境](#开发环境)
  - [编辑器](#编辑器)
  - [测试](#测试)
  - [编译](#编译)

</td></tr>
</table>

## 安装

如果你已安装了 `node.js`，那么可以运行以下命令进行安装

```bash
npm install @axetroy/watchdog -g
```

如果你是用 `Linux`/`MacOS` 系统，可以运行以下命令进行安装

```shell
# 安装最新版本
curl -fsSL https://raw.githubusercontent.com/axetroy/watchdog/master/install.sh | bash
# 安装指定版本
curl -fsSL https://raw.githubusercontent.com/axetroy/watchdog/master/install.sh | bash -s v1.3.0
# 从 gobinaries.com 中安装
curl -sf https://gobinaries.com/axetroy/watchdog@v1.3.0 | sh
```

或者你可以通过[下载页面](https://github.com/axetroy/watchdog/releases)下载你对应平台的可执行文件

> 注意: watchdog 只有一个可执行文件，没有任何其他依赖

最后确保 `watchdog` 可以正常运行

```bash
watchdog --help
```

## 快速开始

首先先创建一个配置文件 `watchdog.config.json`

```json
{
  "service": [
    {
      "name": "本地服务",
      "protocol": "http",
      "addr": "http://127.0.0.1:1080",
      "interval": 10
    }
  ]
}
```

然后运行以下命令启动服务

```bash
# 读取配置，然后监听 9999 端口
watchdog --config=./watchdog.config.json --port=9999
```

服务启动成功，用浏览器打开 `http://localhost:9999` 页面查看监控状态

## FAQ

### 工作原理

看门狗使用协程，为每个服务定时检测，然后将服务的结果通过 Websocket 推送到前端，最后根据服务的配置协议，推送到不同的对象当中。

### 路线图

致力于打造一个简单，易用，可扩展的工具，零依赖，无数据库。

### 常见问题

Q: 会支持数据库数据持久化吗？

A: 不会

## 配置文件

看门狗的配置文件是一个 JSON5 文件，可以在 JSON 文件中写入注释，通常命名为 `watchdog.config.json`

| 字段     | 类型                   | 必填 | 说明                                |
| -------- | ---------------------- | ---- | ----------------------------------- |
| interval | int                    |      | 每个服务检测的间隔时间，单位为 `秒` |
| service  | [[]Service](#服务配置) | \*   | 服务配置                            |

## 服务配置

| 字段     | 类型                          | 必填 | 说明                                                        |
| -------- | ----------------------------- | ---- | ----------------------------------------------------------- |
| name     | string                        | \*   | 服务名称，并且唯一                                          |
| protocol | [string](#支持的协议)         | \*   | 服务协议                                                    |
| addr     | string                        | \*   | 服务地址                                                    |
| interval | int                           | \*   | 服务检测的间隔时间，单位为 `秒`，如果不设置，则使用全局配置 |
| reporter | [[]Reporter](#支持的通知渠道) | \*   | 服务状态变更的通知渠道                                      |

## 支持的协议

| 协议  | 说明                    | addr 字段                                             |
| ----- | ----------------------- | ----------------------------------------------------- |
| ftp   | 检测 FTP 服务           | `FTP` 协议的地址，例如 `ftp://localhost:22`           |
| sftp  | 检测 SFTP 服务          | `SFTP` 协议的地址，例如 `sftp://localhost:22`         |
| http  | 检测 HTTP 服务          | `HTTP` 协议的地址，例如 `http://localhost:22`         |
| https | 检测 HTTPS 服务         | `HTTPS` 协议的地址，例如 `https://localhost:22`       |
| nfs   | -                       | -                                                     |
| pop3  | -                       | -                                                     |
| smb   | -                       | -                                                     |
| smtp  | -                       | -                                                     |
| ssh   | -                       | -                                                     |
| tcp   | 检测 TCP 服务           | `TCP` 协议的地址，例如 `localhost:22`                 |
| udp   | 检测 UDP 服务           | `UDP` 协议的地址，例如 `localhost:22`                 |
| ws    | 检测 WebSocket 服务     | `WebSocket` 协议的地址，例如 `ws://localhost:22`      |
| wss   | 检测 WebSocket SSL 服务 | `WebSocket SSL` 协议的地址，例如 `wss://localhost:22` |

## 支持的通知渠道

| 字段     | 类型                | 必填 | 说明                                     |
| -------- | ------------------- | ---- | ---------------------------------------- |
| protocol | [string](#通知协议) | \*   | 通知协议                                 |
| target   | []string            | \*   | 通知的目标，根据协议不同，填写的内容不同 |
| payload  | interface{}         |      | 针对不同的通知渠道所设置定的配置字段     |

通知渠道的 `protocol` 字段

### 通知协议

| 协议                        | 说明                      |
| --------------------------- | ------------------------- |
| [webhook](#Webhook)         | 通过调用 Webhook 进行通知 |
| [pop3](#POP3)               | 通过 POP3 协议发送邮件    |
| [smtp](#SMTP)               | 通过 SMTP 协议发送邮件    |
| [wechat](#wechat)           | 微信的推送                |
| [wechat-work](#wechat-work) | 企业微信的推送            |

#### Webhook

Webhook 通道即调用 `HTTP` 的 `POST` 方法，请求目标地址，由目标服务器处理

它将向目标 URL 发送 POST 请求，以 `application/json` 类似发送如下字段

| 字段    | 类型   | 说明                                       |
| ------- | ------ | ------------------------------------------ |
| content | string | 消息内容，这是由看门狗拼接而成的消息字符串 |

例如

```json
{
  "reporter": [
    {
      "protocol": "webhook",
      "target": ["https://example.com"]
    }
  ]
}
```

#### POP#

TODO

#### smtp

TODO

#### wechat

微信通道使用的是第三方的微信公众号推送，所以需要第三方服务。

这里选用 [wxPusher](https://wxpusher.zjiecode.com)

payload 是一个 `key-value` 的字典对象，字段如下：

| 字段      | 类型   | 必填 | 说明                                                  |
| --------- | ------ | ---- | ----------------------------------------------------- |
| app_token | string | \*   | [wxPusher](https://wxpusher.zjiecode.com) 的 appToken |

例如

```json
{
  "reporter": [
    {
      "protocol": "wechat",
      "target": ["uid1", "uid2"],
      "payload": {
        "app_token": "xxxxx"
      }
    }
  ]
}
```

#### wechat-work

TODO

## 使用例子

### 单服务

```jsonc
{
  "service": [
    {
      "name": "主站点",
      "protocol": "https",
      "addr": "https://example.com",
      "interval": 10
    }
  ]
}
```

### 多服务

```jsonc
{
  "service": [
    {
      "name": "主站点",
      "protocol": "https",
      "addr": "https://example.com",
      "interval": 10
    },
    {
      "name": "通知 Socket",
      "protocol": "wss",
      "addr": "wss://example.com/api/ws",
      "interval": 10
    }
  ]
}
```

#### 通知渠道

#### 微信通知

```jsonc
{
  "service": [
    {
      "name": "主站点",
      "protocol": "https",
      "addr": "https://example.com",
      "interval": 10
      "reporter": [
        {
          "protocol": "wechat",
          "target": ["user_id_1"],
          "payload": {
            "app_token": "xxxxxx"
          }
        }
      ]
    },
  ]
}
```

#### Webhook 通知

```jsonc
{
  "service": [
    {
      "name": "主站点",
      "protocol": "https",
      "addr": "https://example.com",
      "interval": 10
      "reporter": [
        {
          "protocol": "webhook",
          "target": ["https://example.com/api/webhook/error"]
        }
      ]
    },
  ]
}
```

#### 邮件通知

TODO
