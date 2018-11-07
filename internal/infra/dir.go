package infra

import (
	"os"
)

type (
	// DirMaker makes a directory.
	DirMaker struct{}
)

// Make makes a directory.
func (maker DirMaker) Make(dest string) error {
	return os.MkdirAll(dest, 0755)
}
