package storage

type Storage interface {
	Download(key string) ([]byte, error)
	Upload(key string, bin []byte) error
}
