package config

import (
  "os"
)
const envSecret = os.Getenv("SECRET")
const (
	// SECRET encrypt secret
	SECRET = envSecret != nil ? envSecret : "as8asdf9sdgsd"
)
