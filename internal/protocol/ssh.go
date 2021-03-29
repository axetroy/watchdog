package protocol

import (
	"context"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	knownhosts "golang.org/x/crypto/ssh/knownhosts"
)

func dialSSH(ctx context.Context, network, addr string, config *ssh.ClientConfig) (*ssh.Client, error) {
	d := net.Dialer{Timeout: config.Timeout}
	conn, err := d.DialContext(ctx, network, addr)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	c, chans, reqs, err := ssh.NewClientConn(conn, addr, config)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return ssh.NewClient(c, chans, reqs), nil
}

func PingSSH(ctx context.Context, addr string, auth interface{}) error {
	type AuthWithPassword struct {
		Username string `json:"username"` // 用户名
		Password string `json:"password"` // 密码
	}

	type AuthWithPrivateKey struct {
		Username           string  `json:"username"`             // 用户名
		PrivateKey         string  `json:"private_key"`          // 私钥
		KnownHostsFilepath *string `json:"known_hosts_filepath"` // KnownHosts 文件路径
	}

	config := ssh.ClientConfig{
		Timeout: time.Second * 30,
	}

	if auth != nil {
		authPassword := AuthWithPassword{}
		authPrivateKey := AuthWithPrivateKey{}
		if err := mapstructure.Decode(auth, &authPassword); err == nil {
			// 用户名 + 密码
			config.User = authPassword.Username
			config.Auth = []ssh.AuthMethod{ssh.Password(authPassword.Password)}
			config.HostKeyCallback = ssh.InsecureIgnoreHostKey()
		} else if err := mapstructure.Decode(auth, &authPrivateKey); err == nil {
			// 用户名 + 私钥
			signer, err := ssh.ParsePrivateKey([]byte(authPrivateKey.PrivateKey))

			if err != nil {
				return errors.WithStack(err)
			}

			homeDir, err := os.UserHomeDir()

			if err != nil {
				return errors.WithStack(err)
			}

			if authPrivateKey.KnownHostsFilepath == nil {
				knownHostsFilepath := filepath.Join(homeDir, ".ssh", "known_hosts")

				if _, err = os.Stat(knownHostsFilepath); err == nil {
					hostKeyCallback, err := knownhosts.New(knownHostsFilepath)

					if err != nil {
						return errors.WithStack(err)
					}

					config.HostKeyCallback = hostKeyCallback

				} else if os.IsNotExist(err) {
					// do nothing
				} else {
					return errors.WithStack(err)
				}
			} else {
				hostKeyCallback, err := knownhosts.New(*authPrivateKey.KnownHostsFilepath)

				if err != nil {
					return errors.WithStack(err)
				}

				config.HostKeyCallback = hostKeyCallback
			}

			config.User = authPrivateKey.Username
			config.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
		} else {
			return errors.New("invalid auth for ssh protocol")
		}
	} else {
		config.HostKeyCallback = ssh.InsecureIgnoreHostKey()
	}

	client, err := dialSSH(ctx, "tcp", addr, &config)

	if err != nil {
		if auth == nil && strings.HasPrefix(err.Error(), "ssh: handshake failed:") {
			return nil
		}
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
