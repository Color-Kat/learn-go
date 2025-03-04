package files

import (
	"fmt"
	"os"
)

func WriteFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}(file)

	_, err = file.Write(data)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}
	fmt.Println("File written successfully")

	return nil
}

func ReadFile(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}
	return file, nil
}
