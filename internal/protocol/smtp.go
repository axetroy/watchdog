package protocol

import (
	"net/smtp"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

type loginAuth struct {
	username, password string
}

func NewLoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown fromServer")
		}
	}
	return nil, nil
}

type AuthWithPassword struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func PingSMTP(addr string, auth interface{}) error {

	client, err := smtp.Dial(addr)

	if err != nil {
		return errors.WithStack(err)
	}

	defer client.Close()

	if auth != nil {
		authPassword := AuthWithPassword{}
		if err := mapstructure.Decode(auth, &authPassword); err == nil {
			client, err := smtp.Dial(addr)

			if err != nil {
				return errors.WithStack(err)
			}

			if err = client.Auth(NewLoginAuth(authPassword.Username, authPassword.Password)); err != nil {
				return err
			}
		}
	}

	return nil
}
