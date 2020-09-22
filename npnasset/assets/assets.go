package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"mime"
	"path"
	"path/filepath"

	"emperror.dev/errors"
)

func Asset(base string, p string) ([]byte, string, string, error) {
	var b bytes.Buffer

	file := path.Join(base, p)
	println("!!!!")
	println(file)
	println(filepath.Abs(file))

	data, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, "", "", errors.WithStack(errors.Wrap(err, "error reading asset at ["+p+"]"))
	}

	if data != nil {
		w := gzip.NewWriter(&b)
		_, _ = w.Write(data)
		_ = w.Close()
		data = b.Bytes()
	}

	sum := md5.Sum(data)

	return data, hex.EncodeToString(sum[1:]), mime.TypeByExtension(filepath.Ext(file)), nil
}
