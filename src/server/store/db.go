package store

import (
	"io/ioutil"
)

const dbPath = "storage.db"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Put(path string) {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		// if not existed, then create
		file, e := os.Create(dbPath)
		if e {
			panic(e)
		}
	}
}
