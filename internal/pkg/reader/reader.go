package readerPkg

import "io"

func Read(reader io.Reader) []byte {
	data := make([]byte, 0, 4096)
	tmp := make([]byte, 1)
	var err error
	for {
		_, err = reader.Read(tmp)
		if err != nil {
			break
		}
		data = append(data, tmp...)
	}
	return data
}
