//go:generate go-syncmap -type "states<string,stateObj>"

package openid

import (
	"sync"
	"time"
)

// states хранит сессии авторизации.
type states sync.Map

// stateObj используется для хранения состояний между началом авторизации и
// получением ответа от сервера авторизации.
type stateObj struct {
	Created     time.Time   // время создания
	RedirectURI string      // адрес для редиректа после авторизации
	Data        interface{} // дополнительные данные
}
