package ssh

import (
	"fmt"
	"io"
	"time"

	"golang.org/x/crypto/ssh"
)

type Client struct {
	config *ssh.ClientConfig
	client *ssh.Client
}

type Session struct {
	session *ssh.Session
}

func NewClient(host, username, password string, timeout time.Duration) (*Client, error) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         timeout * time.Second,
	}

	return &Client{
		config: config,
	}, nil
}

func (c *Client) Connect(host string) error {
	client, err := ssh.Dial("tcp", host+":22", c.config)
	if err != nil {
		return fmt.Errorf("SSH bağlantısı kurulamadı: %w", err)
	}

	c.client = client
	return nil
}

func (c *Client) ExecuteCommand(command string) (string, error) {
	if c.client == nil {
		return "", fmt.Errorf("SSH bağlantısı kurulmamış")
	}

	session, err := c.client.NewSession()
	if err != nil {
		return "", fmt.Errorf("SSH session oluşturulamadı: %w", err)
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", fmt.Errorf("komut çalıştırılamadı: %w", err)
	}

	return string(output), nil
}

func (c *Client) Close() error {
	if c.client != nil {
		return c.client.Close()
	}
	return nil
}

func (c *Client) IsConnected() bool {
	return c.client != nil
}
