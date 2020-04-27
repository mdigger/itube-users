package main

import (
	"flag"
	"fmt"
	"itube/users/pkg/email"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/matcornic/hermes/v2"
	yaml "gopkg.in/yaml.v2"
)

// Domain группирует формат писем по доменам.
type Domain struct {
	hermes.Product `yaml:",inline"`       // общее описание домента
	Emails         map[string]hermes.Body `yaml:"templates"`
}

// Config описывает формат писем.
type Config map[string]Domain // список поддерживаемых доменов с конфигурациями

// Generate генерирует общий шаблон со всем шаблонами писем для всех доменов,
// которые определены в конфигурации.
func (cfg Config) Generate() (map[string]email.Domain, error) {
	var result = make(map[string]email.Domain, len(cfg))
	for name, domain := range cfg {
		var h = hermes.Hermes{Product: domain.Product}
		var emails = make(map[string]email.Template, len(domain.Emails))
		for name, mail := range domain.Emails {
			var e = hermes.Email{Body: mail}
			html, err := h.GenerateHTML(e)
			if err != nil {
				return nil, err
			}
			html = regexp.
				MustCompile(`(?:\n\s*)+`).
				ReplaceAllString(html, "\n")
			text, err := h.GeneratePlainText(e)
			if err != nil {
				return nil, err
			}
			var subject string
			switch name {
			case "EMAIL":
				subject = "Confirm your account"
			case "PASSWORD":
				subject = "Reset your password"
			default:
				subject = fmt.Sprintf("Unknown subject for %q", name)
			}
			emails[name] = email.Template{
				Subject: subject,
				Text:    text,
				HTML:    html,
			}
		}
		purk, err := url.Parse(domain.Link)
		if err != nil {
			return nil, err
		}
		result[name] = email.Domain{
			From:   fmt.Sprintf("noreply@%s", strings.ToLower(purk.Host)),
			Emails: emails,
		}
	}
	return result, nil
}

// Load загружает конфигурационный файл и разбирает его
func Load(filename string) (*Config, error) {
	file, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return nil, err
	}
	var config = new(Config)
	err = yaml.NewDecoder(file).Decode(config)
	_ = file.Close()
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return config, nil
}

func main() {
	config := flag.String("config", "email_config.yaml", "configuration file")
	output := flag.String("out", "email_templates.yaml", "output email templates file")
	flag.Parse()
	log.SetFlags(0)

	log.Println("loading config", *config)
	cfg, err := Load(*config)
	if err != nil {
		log.Fatal(err)
	}
	result, err := cfg.Generate()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("saving generated template to file", *output)
	file, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}
	var enc = yaml.NewEncoder(file)
	err = enc.Encode(result)
	_ = enc.Close()
	_ = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("template generated")
}
