package compress

import (
	"bytes"
	"compress/zlib"
	"github.com/foobaz/go-zopfli/zopfli"
	"io/ioutil"
)

func Compress(s []byte) ([]byte, error) {
	o := zopfli.DefaultOptions()
	w := bytes.NewBuffer([]byte{})

	if err := zopfli.ZlibCompress(&o, s, w); err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}

func Decompress(b []byte) ([]byte, error) {
	r, err := zlib.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(r)
}
