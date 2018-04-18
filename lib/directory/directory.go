package directory

import (
	"io/ioutil"

	"github.com/flexphere/sts/settings"
)

type Directory struct {
}

func New() *Directory {
	return &Directory{}
}

func (d *Directory) Download(key string) ([]byte, error) {
	return ioutil.ReadFile(*settings.Settings.TMP_DIRECTORY + key)
}

func (d *Directory) Upload(key string, bin []byte) error {
	return ioutil.WriteFile(*settings.Settings.TMP_DIRECTORY+key, bin, 0644)
}
