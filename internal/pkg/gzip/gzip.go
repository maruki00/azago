package gzipPkg

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

func Decompress(body []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	decompressedData, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("error reading decompressed data: %s", err.Error())
	}
	return decompressedData, nil
}

func Compress(body []byte) ([]byte, error) {
	var buff bytes.Buffer
	writer := gzip.NewWriter(&buff)
	defer writer.Close()
	if _, err := writer.Write(body); err != nil {
		return []byte{}, fmt.Errorf("could not write data to buffer: %v", err)
	}
	return buff.Bytes(), nil
}
