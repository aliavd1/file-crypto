package fileutils

import (
	"os"
)

func WriteToFile(data []byte, file string) {
	if err := os.WriteFile(file, data, 777); err != nil {
		panic(err.Error())
	}
}

func ReadFromFile(file string) []byte {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err.Error())
	}
	return data
}
