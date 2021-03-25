package protocol

import (
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

func PingSSH(addr string, auth interface{}) error {
	type AuthWithPassword struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type AuthWithPrivateKey struct {
		Username   string `json:"username"`
		PrivateKey string `json:"private_key"`
	}

	config := ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 60,
	}

	if auth != nil {
		authPassword := AuthWithPassword{}
		authPrivateKey := AuthWithPrivateKey{}
		if err := mapstructure.Decode(auth, &authPassword); err == nil {
			// 用户名 + 密码
			config.User = authPassword.Username
			config.Auth = []ssh.AuthMethod{ssh.Password(authPassword.Password)}
		} else if err := mapstructure.Decode(auth, &authPrivateKey); err == nil {
			// 用户名 + 私钥
			signer, err := ssh.ParsePrivateKey([]byte(authPrivateKey.PrivateKey))

			if err != nil {
				return errors.WithStack(err)
			}

			config.User = authPrivateKey.Username
			config.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
		}
	}

	client, err := ssh.Dial("tcp", addr, &config)

	if err != nil {
		return errors.WithStack(err)
	}

	defer client.Close()

	session, err := client.NewSession()

	if err != nil {
		return errors.WithStack(err)
	}

	defer session.Close()

	return nil
}
