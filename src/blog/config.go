package main

func ReadConfig(path string, byteLen int) (map[string]interface{}, error) {
	f, _ := os.Open(path)
	defer f.Close()

	m := make(map[string]interface{})

	b := make([]byte, byteLen)
	length, readErr := f.Read(b)
	if readErr != nil {
		return m, readErr
	}
	filebyte := b[:length]
	err := json.Unmarshal(filebyte, &m)
	return m, err
}
