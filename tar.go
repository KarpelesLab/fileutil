package fileutil

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// TarExtract will extract the given input as tar data to the specified directory
//
// This method will not check for anything nefarious. Only use for trusted archives.
func TarExtract(r io.Reader, targetDir string) error {
	arch := tar.NewReader(r)

	for {
		h, err := arch.Next()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("while extracting: %w", err)
		}
		tgt := filepath.Join(targetDir, h.Name)
		fi := h.FileInfo()
		switch fi.Mode().Type() {
		case os.ModeDir:
			err = os.MkdirAll(tgt, fi.Mode())
			if err != nil {
				return fmt.Errorf("while creating dir %s: %w", h.Name, err)
			}
			os.Chmod(tgt, fi.Mode()&0755) // force 0755 max
		case os.ModeSymlink:
			err = os.Symlink(tgt, h.Linkname)
			if err != nil {
				return fmt.Errorf("while creating symlink %s: %w", h.Name, err)
			}
		default:
			if fi.Mode().Type().IsRegular() {
				os.MkdirAll(filepath.Dir(tgt), 0755) // ensure dir first
				err = WriteFileReader(tgt, arch, fi.Mode()&0755)
				if err != nil {
					return fmt.Errorf("while creating file %s: %w", h.Name, err)
				}
				break
			}
			return fmt.Errorf("File %s has unsupported type %s", h.Name, fi.Mode().Type())
		}
	}
}
