package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config struct
type Config struct {
	ID             int
	Username       string
	Salt           []byte
	HashedPassword string
	Role           string
}

// Update update config from other
func (c *Config) Update(s Config) {
	if s.Username != "" {
		c.Username = s.Username
	}
	if s.Salt != nil {
		c.Salt = s.Salt
	}
	if s.HashedPassword != "" {
		c.HashedPassword = s.HashedPassword
	}
}

const dbPath = "storage.db"

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

// Update config
func Update(config Config) {
	_, err := os.Stat(dbPath)
	if err != nil && os.IsNotExist(err) {
		// file not existed
		fmt.Println("db file not existed")
	} else {
		c := Read()
		c.Update(config)
	}
	s, _ := json.Marshal(config)
	e := ioutil.WriteFile(dbPath, s, 0644)
	check(e)
}

// Read config
func Read() Config {
	raw, err := ioutil.ReadFile(dbPath)
	check(err)
	config := Config{}
	json.Unmarshal(raw, &config)
	return config
}
