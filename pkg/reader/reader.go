package readerPkg

import (
	"bufio"
	"bytes"
	"io"
)

func Read(reader io.Reader) ([]byte, error) {

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ReadUntil(r io.Reader, del []byte) ([]byte, error) {
	reader := bufio.NewReader(r)
	var buffer bytes.Buffer

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			return nil, err
		}
		buffer.Write(line)

		if bytes.HasSuffix(buffer.Bytes(), del) {
			break
		}
	}
	return buffer.Bytes(), nil
}
