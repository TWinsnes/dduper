package dduper

import (
	"testing"

	"github.com/spf13/afero"
)

func TestDeleteDuplicates(t *testing.T) {
	fsys := afero.NewMemMapFs()
	fsys.Mkdir("dir1", 0644)
	fsys.Mkdir("dir2", 0644)
	afero.WriteFile(fsys, "file1.txt", []byte("content"), 0644)
	afero.WriteFile(fsys, "dir1/file2.txt", []byte("content"), 0644)
	afero.WriteFile(fsys, "dir2/file1.txt", []byte("notdupe"), 0644)

	result, _ := ScanForDupes(fsys, ".")

	deleteDuplicates(fsys, result.DupePaths)

	if b, _ := afero.Exists(fsys, "file1.txt") ; b {
		t.Error("The file 'file1.txt' was not removed")
	}

}

func TestScanForDupes(t *testing.T) {
	var wantFilesScanned = 3
	var wantDupes = 1

	fsys := afero.NewMemMapFs()
	fsys.Mkdir("dir1", 0644)
	fsys.Mkdir("dir2", 0644)
	afero.WriteFile(fsys, "file1.txt", []byte("content"), 0644)
	afero.WriteFile(fsys, "dir1/file2.txt", []byte("content"), 0644)
	afero.WriteFile(fsys, "dir2/file1.txt", []byte("notdupe"), 0644)

	result, err := ScanForDupes(fsys, ".")

	if err != nil {
		t.Errorf("Got error: %s", err)
	}

	if result.FilesScanned != wantFilesScanned {
		t.Errorf("Expected files scanned to be %d, got %d", wantFilesScanned, result.FilesScanned)
	}

	if len(result.DupePaths) != wantDupes {
		t.Errorf("Expected %d dupes, got %d", wantDupes, len(result.DupePaths))
	}
}