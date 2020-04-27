package email

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"gopkg.in/mail.v2"
	"gopkg.in/yaml.v2"
)

// Templates описывает список доменов и поддерживаемые для них шаблоны писем.
type Templates struct {
	*mail.Dialer
	list map[string]Domain
}

// Init загружает шаблоны писем из файла в формате yaml.
func Init(connection, templatesFilename string) (*Templates, error) {
	// инициализируем настройка для подключения к SMTP
	dialer, err := Dialer(connection)
	if err != nil {
		return nil, err
	}
	// загружаем и разбираем файл с шаблонами писем
	file, err := os.Open(filepath.Clean(templatesFilename))
	if err != nil {
		return nil, err
	}
	var tmplts = make(map[string]Domain)
	err = yaml.NewDecoder(file).Decode(&tmplts)
	_ = file.Close()
	if err != nil {
		return nil, err
	}
	return &Templates{
		Dialer: dialer,
		list:   tmplts,
	}, nil
}

// Domains возвращает список поддерживаемых доменов.
func (t Templates) Domains() []string {
	var result = make([]string, 0, len(t.list))
	for name := range t.list {
		result = append(result, name)
	}
	sort.Strings(result)
	return result
}

// Domain возвращает почтовые шаблоны для указанного домента.
func (t Templates) Domain(name string) (*Domain, error) {
	domain, ok := t.list[name]
	if !ok {
		return nil, fmt.Errorf("unsupported domain %q", name)
	}
	return &domain, nil
}

// Domain описывает конфигурацию для домена.
type Domain struct {
	From   string              // от кого отправляется письмо
	Emails map[string]Template // список поддерживаемых типов писем
}

// Email возвращает почтовые шаблоны для указанного домента.
func (d Domain) Email(name string) (*Template, error) {
	email, ok := d.Emails[name]
	if !ok {
		return nil, fmt.Errorf("unsupported email template %q", name)
	}
	return &email, nil
}
