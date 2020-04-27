package email

import (
	"fmt"
	"net/url"
	"strconv"

	"gopkg.in/mail.v2"
)

// Dialer инициализируем объект для отправки почты.
func Dialer(connection string) (*mail.Dialer, error) {
	// разбираем параметры подключения к серверу SMTP
	parsed, err := url.Parse(connection)
	if err != nil {
		return nil, err
	}
	var scheme = parsed.Scheme
	var port int
	switch scheme {
	case "smtp":
		port = 587
	case "smtps":
		port = 465
	default:
		return nil, fmt.Errorf("bad smtp url scheme: %v", scheme)
	}
	if p := parsed.Port(); p != "" {
		port, err = strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
	}
	var hostname = parsed.Hostname()
	if hostname == "" {
		hostname = "localhost"
	}
	var username = parsed.User.Username()
	var password, _ = parsed.User.Password()

	var dialer = mail.NewDialer(hostname, port, username, password)
	dialer.StartTLSPolicy = mail.MandatoryStartTLS
	return dialer, nil
}
