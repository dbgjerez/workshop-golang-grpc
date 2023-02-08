package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	c "github.com/dbgjerez/workshop-golang-grpc/comment/grpc"
)

const (
	ENV_CONFIG_JSON_LOCATION               string = "ENV_CONFIG_JSON_LOCATION"
	ENV_CONFIG_JSON_LOCATION_DEFAULT_VALUE string = "./data.json"
)

type Store struct {
	Path string
}

func NewStore() *Store {
	path := GetEnv(ENV_CONFIG_JSON_LOCATION,
		ENV_CONFIG_JSON_LOCATION_DEFAULT_VALUE)
	return &Store{Path: path}
}

func (store *Store) Exists() bool {
	if _, err := os.Stat(store.Path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

// TODO read before the file and store the content in memory

func (store *Store) ReadStore() ([]*c.Comment, error) {
	f, err := ioutil.ReadFile(store.Path)
	if err != nil {
		return []*c.Comment{}, errors.New(fmt.Sprintf("Error reading the file %s: %v", store.Path, err))
	}
	var comments []*c.Comment
	err = json.Unmarshal(f, &comments)
	if err != nil {
		return []*c.Comment{}, errors.New(fmt.Sprintf("Error unmarshal the %s content: %v", store.Path, err))
	}
	return comments, nil
}

func GetEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
