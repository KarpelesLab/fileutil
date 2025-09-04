package fileutil

import (
	"bytes"
	"os"
	"path/filepath"
)

// Put atomically writes data to a file if it doesn't already contain the same content.
// It returns true if the file was written, false if it already contained the data.
// The operation is atomic - the file will either be fully written or not modified at all.
func Put(filename string, data []byte, perm os.FileMode) (bool, error) {
	// Check if target file already contains identical data
	if st, err := os.Stat(filename); err == nil {
		if st.Size() == int64(len(data)) {
			// Only read file if size matches (optimization)
			if old, err := os.ReadFile(filename); err == nil {
				if bytes.Equal(old, data) {
					// Ensure permissions are correct
					if st.Mode().Perm() != perm {
						return false, os.Chmod(filename, perm)
					}
					return false, nil
				}
			}
		}
	}

	// Create temporary file in the same directory for atomic rename
	f, err := os.CreateTemp(filepath.Dir(filename), ".tmp-")
	if err != nil {
		return false, err
	}
	tempName := f.Name()
	defer os.Remove(tempName) // Clean up on any error
	defer f.Close()

	// Write data
	if _, err = f.Write(data); err != nil {
		return false, err
	}
	
	// Set permissions
	if err = f.Chmod(perm); err != nil {
		return false, err
	}
	
	// Close before rename (required on some systems)
	if err = f.Close(); err != nil {
		return false, err
	}
	
	// Atomic rename
	return true, os.Rename(tempName, filename)
}