package infra

import (
	"io/ioutil"
	"os"
)

type (
	// FileReader reads a file.
	FileReader struct{}
	// FileWriter writes data to a file.
	FileWriter struct{}
	// FileExist checks whether a file exists.
	FileExist struct{}
)

// Read reads a file.
func (reader FileReader) Read(src string) (string, error) {
	b, err := ioutil.ReadFile(src)
	return string(b), err
}

// Write writes data to a file.
func (writer FileWriter) Write(dest string, content []byte) error {
	return ioutil.WriteFile(dest, content, 0644)
}

// Exist returns whether a file exists.
func (exist FileExist) Exist(src string) bool {
	_, err := os.Stat(src)
	return err == nil
}
