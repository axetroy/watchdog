package protocol

import (
	"time"

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
		if entry, ok := auth.(AuthWithPassword); ok {
			// 用户名 + 密码
			config.User = entry.Username
			config.Auth = []ssh.AuthMethod{ssh.Password(entry.Password)}
		} else if entry, ok := auth.(AuthWithPrivateKey); ok {
			// 用户名 + 私钥
			signer, err := ssh.ParsePrivateKey([]byte(entry.PrivateKey))

			if err != nil {
				return errors.WithStack(err)
			}

			config.User = entry.Username
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
