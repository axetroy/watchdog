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
	Interval uint      `json:"interval" validate:"required,gt=0"` // 检测间隔时间，单位秒
	Service  []Service `json:"service" validate:"required,dive"`  // 检测的服务
}

type Service struct {
	Name     string      `json:"name" validate:"required"`                     // 服务名称
	Protocol string      `json:"protocol" validate:"required,detect_protocol"` // 服务协议, 支持 http/https/ws/wss/tcp/tdp/ssh
	Addr     string      `json:"addr" validate:"required"`                     // 地址
	Interval uint        `json:"interval" validate:"gt=0"`                     // 检测任务的间隔时间，单位秒
	Report   []Reporter  `json:"report" validate:"dive"`                       // 通知渠道，支持多个通知渠道
	Auth     interface{} `json:"auth"`                                         // 身份认证的字段，任意类型，主要是每个协议可能需要的字段不一样
}

// 消息通知渠道
type Reporter struct {
	Protocol string      `json:"protocol" validate:"required,notify_protocol"` // 协议，支持 wechat/webhook/smtp
	Target   []string    `json:"target" validate:"required"`                   // 推送的目标，可以是多个目标
	Payload  interface{} `json:"payload"`                                      // 额外的数据，根据不同的协议，所带的数据不同
}

var (
	validate = validator.New()
	trans    ut.Translator
)

func isValidDetectProtocol(protocol string) bool {
	supportDetectProtocols := map[string]bool{
		"http":  true,
		"https": true,
		"ws":    true,
		"wss":   true,
		"tcp":   true,
		"udp":   true,
		"ftp":   true,
		"sftp":  true,
		"ssh":   true,
		"smtp":  true,
		"pop3":  false,
		"nfs":   false,
		"smb":   false,
	}

	if isSsupport, ok := supportDetectProtocols[protocol]; ok {
		return isSsupport
	} else {
		return false
	}
}

func isValidNotifyProtocol(protocol string) bool {
	supportNotifyProtocols := map[string]bool{
		"wechat":  true,
		"webhook": true,
		"smtp":    true,
		"pop3":    false,
	}

	if isSsupport, ok := supportNotifyProtocols[protocol]; ok {
		return isSsupport
	} else {
		return false
	}
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

	if err := validate.RegisterValidation("detect_protocol", func(field validator.FieldLevel) bool {
		val := field.Field().String()

		return isValidDetectProtocol(val)
	}); err != nil {
		panic(err)
	}

	if err := validate.RegisterValidation("notify_protocol", func(field validator.FieldLevel) bool {
		val := field.Field().String()

		return isValidNotifyProtocol(val)
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
