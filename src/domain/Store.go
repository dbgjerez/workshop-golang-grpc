package domain

import (
	"comment/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	ENV_CONFIG_JSON_LOCATION               string = "ENV_CONFIG_JSON_LOCATION"
	ENV_CONFIG_JSON_LOCATION_DEFAULT_VALUE string = "./data.json"
)

type Store struct {
	Path string
}

func NewStore() *Store {
	path := utils.GetEnv(ENV_CONFIG_JSON_LOCATION,
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

func (store *Store) ReadStore() ([]*Comment, error) {
	f, err := ioutil.ReadFile(store.Path)
	if err != nil {
		return []*Comment{}, errors.New(fmt.Sprintf("Error reading the file %s: %v", store.Path, err))
	}
	var comments []*Comment
	err = json.Unmarshal(f, &comments)
	if err != nil {
		return []*Comment{}, errors.New(fmt.Sprintf("Error unmarshal the %s content: %v", store.Path, err))
	}
	return comments, nil
}
