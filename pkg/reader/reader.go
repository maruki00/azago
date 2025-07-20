package readerPkg

import (
	"io"
)

func Read(reader io.Reader) ([]byte, error) {

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return data, nil
}
