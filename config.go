package watchdog

import (
	"io/ioutil"
	"log"
	"reflect"
	"strings"

	"github.com/fatih/color"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/pkg/errors"
	"github.com/yosuke-furukawa/json5/encoding/json5"
)

type Config struct {
	Interval uint      `json:"interval" validate:"required"` // 检测间隔时间
	Service  []Service `json:"service" validate:"required"`  // 检测的服务
}

type Service struct {
	Name         string         `json:"name" validate:"required"`              // 服务名称
	Protocol     string         `json:"protocol" validate:"required,protocol"` // 服务协议, 支持 http/https/ws/wss/tcp/tdp/ssh
	Addr         string         `json:"addr" validate:"required"`              // 地址
	Interval     uint           `json:"interval"`                              // 检测任务的间隔时间
	Notifycation []Notification `json:"notification"`                          // 通知渠道，支持多个通知渠道
}

// 消息通知渠道
type Notification struct {
	Protocol string   // 协议，支持 wechat/email/slack/webhook
	Target   []string // 推送的目标，可以是多个目标
}

var (
	validate = validator.New()
	trans    ut.Translator
)

func isValidProtocol(protocol string) bool {
	switch protocol {
	case "http":
		fallthrough
	case "https":
		fallthrough
	case "tcp":
		fallthrough
	case "udp":
		fallthrough
	case "ws":
		fallthrough
	case "wss":
	case "ftp":
		fallthrough
	case "sftp":
		fallthrough
	case "ssh":
		return true
	default:
		return false
	}

	return false
}

func init() {
	z := en.New()
	uni := ut.New(z, z)
	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ = uni.GetTranslator("zh")
	if err := enTranslations.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Fatalln(err)
	}

	// register function to get tag name from json tags.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := validate.RegisterValidation("protocol", func(field validator.FieldLevel) bool {
		val := field.Field().String()

		return isValidProtocol(val)
	}); err != nil {
		panic(err)
	}
}

func NewConfig(content []byte) (*Config, error) {
	var (
		config = Config{}
		err    error
	)

	if err = json5.Unmarshal(content, &config); err != nil {
		return nil, errors.WithStack(err)
	}

	err = validate.Struct(config)

	if err != nil {
		// translate all error at once
		errs := err.(validator.ValidationErrors)

		errorsMap := errs.Translate(trans)

		msg := []string{}

		for _, e := range errorsMap {
			msg = append(msg, color.RedString("[config]: "+e))
		}

		return nil, errors.New(strings.Join(msg, "\n"))
	}

	return &config, nil
}

func NewConfigFromFile(configFilepath string) (*Config, error) {
	b, err := ioutil.ReadFile(configFilepath)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return NewConfig(b)
}
