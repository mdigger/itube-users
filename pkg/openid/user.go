package openid

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

// UserInfo описывает информацию о пользователе, получаемую от сервера
// авторизации.
//
// Не все поля будут заполнены для разных пользователей. Для получения большей
// информации или обновления уже существующих полей на их актуальную информацию
// можно использовать метод Update.
type UserInfo struct {
	Issuer        string    `json:"-"`                               // URL провайдера, который выдал этот токен
	Subject       string    `json:"-"`                               // уникальный идентификатор пользователя провайдера
	Email         string    `json:"email"`                           // email-адрес пользователя
	Verified      bool      `json:"email_verified"`                  // флаг, что email-адрес подтвержден
	Name          string    `json:"name,omiempty"`                   // полное отображаемое имя
	FamilyName    string    `json:"family_name,omitempty"`           // фамилия
	GivenName     string    `json:"given_name,omitempty"`            // имя
	MiddleName    string    `json:"middle_name,omitempty"`           // отчество или среднее имя
	NickName      string    `json:"nickname,omitempty"`              // прозвище
	Locale        string    `json:"locale,omitempty"`                // BCP 47 формат
	Picture       string    `json:"picture,omitempty"`               // ссылка на аватар пользователя
	Profile       string    `json:"profile,omitempty"`               // ссылка на профиль пользователя
	Website       string    `json:"website,omitempty"`               // ссылка на сайт пользователя
	Gender        string    `json:"gender,omitempty"`                // пол: female, male или другая строка
	Birthdate     string    `json:"birthdate,omitempty"`             // день рождения (YYYY-MM-DD), год может быть 0000, или наоборот - указан только год YYYY
	Zoneinfo      string    `json:"zoneinfo,omitempty"`              // местоположение: Europe/Paris или America/Los_Angeles
	Phone         string    `json:"phone_number,omitempty"`          // телефон: +1 (604) 555-1234;ext=5678
	PhoneVerified bool      `json:"phone_number_verified,omitempty"` // флаг, что телефонный номер подтвержден
	Updated       int64     `json:"updated_at,omitempty"`            // дата обновления
	Expiry        time.Time `json:"-"`                               // до какого времени эта информация считается актуальной

	provider    *oidc.Provider // провайдер авторизации
	oauth2Token *oauth2.Token  // токен авторизации OAuth2
	data        interface{}    // дополнительная информация
}

// Update обращается к серверу провайдера авторизации и запрашивает полный
// профиль пользователя. На основании полученных данных обновляет информацию
// о пользователе.
//
// Если информация заполнена не из токена идентификации, то обновление не будет
// производиться и ошибка не возвращается.
func (u *UserInfo) Update(ctx context.Context) error {
	// не обновляем, если настройки провайдера и токен авторизации недоступны
	if u.provider == nil || u.oauth2Token == nil {
		return nil
	}
	// обращаемся к внешнему сервису авторизации и запрашиваем обновленные
	// данные из профиля пользователя
	if ctx == nil {
		ctx = context.Background()
	}
	userInfo, err := u.provider.UserInfo(ctx, oauth2.StaticTokenSource(u.oauth2Token))
	if err != nil {
		return fmt.Errorf("error getting user profile update: %w", err)
	}
	// обновлям поля с полученной информацией
	u.Subject = userInfo.Subject
	u.Profile = userInfo.Profile
	u.Email = userInfo.Email
	u.Verified = userInfo.EmailVerified
	// заполняем остальные поля
	if err = userInfo.Claims(u); err != nil {
		return fmt.Errorf("error parsing updated user profile: %w", err)
	}
	return nil
}

// IsExpired возвращает true, если информация считается устаревшей.
// Для вычисления используется дата времени жизни идентификационного токена, из
// которого взята информация о пользователе. В противном случае информация
// всегда будет считаться неустаривающей и возвращаться false.
func (u UserInfo) IsExpired() bool {
	return !u.Expiry.IsZero() && time.Since(u.Expiry) > 0
}

// JSON возвращает представление данные о пользователе в виде JSON.
func (u UserInfo) JSON() []byte {
	type userExclude struct {
		UserInfo
		Email    string `json:"email,omitempty"`
		Verified bool   `json:"email_verified,omitempty"`
	}
	// data, _ := json.MarshalIndent(userExclude{UserInfo: u}, "", "\t")
	data, _ := json.Marshal(userExclude{UserInfo: u})
	return data
}

// Data возвращает дополнительные данные, прикрепленные к изначальному запросу
// получения информации о пользовтаеле.
func (u UserInfo) Data() interface{} {
	return u.data
}
