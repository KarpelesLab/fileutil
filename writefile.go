package fileutil

import (
	"io"
	"os"
)

// WriteFileReader is similar to os.WriteFile but using a Reader instead of a []byte
func WriteFileReader(name string, data io.Reader, perm os.FileMode) error {
	f, err := os.OpenFile(name+"~", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer os.Remove(name + "~")

	_, err = io.Copy(f, data)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	if err == nil {
		err = os.Rename(name+"~", name)
	}
	return err
}
