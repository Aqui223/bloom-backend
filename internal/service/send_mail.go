package service

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"time"

	"github.com/slipe-fun/skid-backend/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type TokenFile struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type xoauth2Auth struct {
	username    string
	accessToken string
}

func (a xoauth2Auth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := fmt.Sprintf("user=%s\x01auth=Bearer %s\x01\x01", a.username, a.accessToken)
	return "XOAUTH2", []byte(s), nil
}

func (a xoauth2Auth) Next(fromServer []byte, more bool) ([]byte, error) {
	return nil, nil
}

func XOAUTH2(username, accessToken string) smtp.Auth {
	return xoauth2Auth{username, accessToken}
}

func SendMail(subject string, content string, recipiement string) error {
	cfg := config.LoadConfig("configs/config.yaml")

	oauthConfig := &oauth2.Config{
		ClientID:     cfg.Email.ClientId,
		ClientSecret: cfg.Email.ClientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://mail.google.com/"},
	}

	token := &oauth2.Token{
		AccessToken:  cfg.Email.AccessToken,
		RefreshToken: cfg.Email.RefreshToken,
		Expiry:       time.Now().Add(-time.Hour),
	}

	token, err := oauthConfig.TokenSource(context.Background(), token).Token()
	if err != nil {
		return err
	}

	email := cfg.Email.Email
	auth := XOAUTH2(email, token.AccessToken)

	dialer := &net.Dialer{
		Timeout:   10 * time.Second,
		DualStack: false,
	}

	conn, err := dialer.Dial("tcp4", "smtp.gmail.com:587")
	if err != nil {
		return err
	}
	defer conn.Close()

	c, err := smtp.NewClient(conn, "smtp.gmail.com")
	if err != nil {
		return err
	}
	defer c.Quit()

	if ok, _ := c.Extension("STARTTLS"); ok {
		tlsConfig := &tls.Config{
			ServerName:         "smtp.gmail.com",
			InsecureSkipVerify: false,
		}
		if err = c.StartTLS(tlsConfig); err != nil {
			return err
		}
	} else {
		return errors.New("server does not support STARTTLS")
	}

	if err = c.Auth(auth); err != nil {
		return err
	}

	if err = c.Mail(email); err != nil {
		return err
	}

	if err = c.Rcpt(recipiement); err != nil {
		return err
	}

	wc, err := c.Data()
	if err != nil {
		return err
	}

	msg := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, content))
	wc.Write(msg)
	wc.Close()

	return nil
}
