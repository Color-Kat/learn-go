package files

import (
	"fmt"
	"os"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(filename string) *JsonDb {
	return &JsonDb{
		filename: filename,
	}
}

func (db *JsonDb) Write(data []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
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
		return
	}

	fmt.Println("File written successfully")
}

func (db *JsonDb) Read() ([]byte, error) {
	file, err := os.ReadFile(db.filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}
