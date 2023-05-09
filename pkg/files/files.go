package files

import (
	"bufio"
	"bytes"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var buffer bytes.Buffer
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		buffer.WriteString(scanner.Text())
	}

	if scanner.Err() != nil {
		return "", scanner.Err()
	}

	return buffer.String(), nil
}
