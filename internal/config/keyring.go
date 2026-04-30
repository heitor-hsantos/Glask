package config

import (
	"github.com/zalando/go-keyring"
	_ "github.com/zalando/go-keyring"
)

const (
	keyringService = "glask"
	keyringUser    = "gemini_api_key"
)

func SaveAPIKey(apiKey string) error {
	return keyring.Set(keyringService, keyringUser, apiKey)
}
func GetAPIKey() (string, error) { return keyring.Get(keyringService, keyringUser) }

func DeleteAPIKey() error { return keyring.Delete(keyringService, keyringUser) }
