package dduper

import (
	"testing"
	"testing/fstest"
)

func TestHashFileConsistent(t *testing.T) {
	fsys1 := fstest.MapFS{"file1": {Data: []byte("content")}}
	fsys2 := fstest.MapFS{"file2": {Data: []byte("content")}}

	f1, _ := fsys1.Open("file1")

	defer f1.Close()

	hash, err := HashFile(f1)

	if err != nil {
		t.Errorf("File 1 hash failed with error: %s", err)
	}

	f2, _ := fsys2.Open("file2")

	defer f1.Close()

	hash2, err := HashFile(f2)

	if err != nil {
		t.Errorf("File 2 hash failed with error: %s", err)
	}

	if hash != hash2 {
		t.Error("Hash for file 1 and file 2 did not match")
	}

}
