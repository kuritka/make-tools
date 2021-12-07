package tempstore

import (
	"encoding/json"
	"errors"
	"github.com/kuritka/make-tools/internal/providers/storage"
	"github.com/kuritka/make-tools/internal/utils/zerolog"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)


const storageIDPath = "/tmp/maketools.json"

type EnvStorage struct {
	store storage.KV
}

var log = zerolog.Logger()
func NewStorage() (s *EnvStorage){
	s = &EnvStorage{}
	return
}

func save(s storage.KV) error {
	bytes, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return os.WriteFile(storageIDPath, bytes, 0600 )
}

func load() (storage.KV, error) {
	s := make(storage.KV,0)
	bytes, err := os.ReadFile(storageIDPath)
	if err != nil {
		return s, nil
	}
	err = json.Unmarshal(bytes,&s)
	return s, err
}

func (e *EnvStorage) Get(variable string) (string, error) {
	s,err := load()
	return s[variable], err
}

func (e *EnvStorage) Set(variable, value string) error {
	s, err := load()
	if err != nil {
		return err
	}
	s[variable] = value
	return save(s)
}

func (e *EnvStorage) Clear() {
	err := os.Remove(storageIDPath)
	if errors.Is(err, os.ErrNotExist) {
		return
	}
	kingpin.FatalIfError(err, "can't remove %s", storageIDPath)
}
