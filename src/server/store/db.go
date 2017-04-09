package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Id             int
	Username       string
	HashedPassword string
	Role           string
}

const dbPath = "storage.db"

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func Update(config Config) {
	_, err := os.Stat(dbPath)
	if err != nil && os.IsNotExist(err) {
		// file not existed
		s, _ := json.Marshal(config)
		e = ioutil.WriteFile(dbPath, s, 0644)
		check(e)
	} else {
		c = Read()
		c = config
		s, _ := json.Marshal(c)
		e = ioutil.WriteFile(dbPath, s, 0644)
	}
}

func Read() Config {
	raw, err := ioutil.ReadFile(dbPath)
	check(err)
	config := Config{}
	json.Unmarshal(raw, &config)
	return config
}
