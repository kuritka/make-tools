package storage

type Storage interface {
	Get(variable string) (string, error)
	Set(variable, value string) error
	Clear()
}

type KV map[string]string

