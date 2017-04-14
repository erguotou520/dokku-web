package config

import (
	"os"
)

var (
	// Secret jwt sign secret
	Secret = "asdfasd&sdfa(66)"
)

func init() {
	secret := os.Getenv("SECRET")
	if secret != "" {
		Secret = secret
	}
}
