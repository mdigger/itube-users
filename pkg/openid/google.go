package openid

// GoogleProviderName задает название провайдера Google, используемое по
// умолчанию.
var GoogleProviderName = "google"

// NewGoogle возвращает инициализированный Auth для работы с Google OpenID
// Connect. Вызывает панику, если сервер https://accounts.google.com не доступен.
func NewGoogle(clientID, secret string) (*Provider, error) {
	provider, err := New(Config{
		Name:     GoogleProviderName,
		URL:      "https://accounts.google.com",
		СlientID: clientID,
		Secret:   secret,
	})
	if err != nil {
		return nil, err
	}
	return provider, nil
}
