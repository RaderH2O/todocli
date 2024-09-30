package fileoperations

import (
	"fmt"
	"os"
)

func WriteToFile(filepath string, object fmt.Stringer) error {
	content := object.String()
	err := os.WriteFile(filepath, []byte(content), 0666)
	return err
}

func ReadFromFile(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	return string(data), err
}
