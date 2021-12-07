package envstorage

import (
	"encoding/json"
	"github.com/kuritka/make-tools/internal/providers/storage"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

const storageID = "ENV_MAKE_TOOLS_STORAGE"

type EnvStorage struct {
	store storage.KV
}

func NewStorage() *EnvStorage {
	if _, b := os.LookupEnv(storageID); !b {
		kingpin.Fatalf("Storage is not initialized. Put 'export %s={}' to the beginning of the makefile", storageID)
	}
	return &EnvStorage{}
}

func save(s storage.KV) error {
	bytes, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return os.Setenv(storageID, string(bytes))
}

func load() (storage.KV, error) {
	s := make(storage.KV,0)
	data := os.Getenv(storageID)
	err := json.Unmarshal([]byte(data),&s)
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
	// no matter if env variable exists or not
	_ = os.Unsetenv(storageID)
}
