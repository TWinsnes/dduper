package dduper

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/fs"
)

// HashFile generates a sha1 hash of a file
func HashFile(f fs.File) (string, error) {

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", nil
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
