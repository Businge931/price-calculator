package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.New("failed to open file")
	}
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("failed to read line in file")

	}
	file.Close()
	return lines, nil
}

func WriteJSON(path string, data any) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to convert data to JSON")
	}
	file.Close()
	return nil
}
