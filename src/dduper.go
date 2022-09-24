package dduper

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/spf13/afero"
)

type DupeScanResult struct {
	FilesScanned int
	DupePaths    []string // list of files that are duplicates
}


// DDupeCurrentDir wraps DDupe and passes in the current working directory.
// If you need to modify the directory scanned, use DDupe
func DDupeCurrentDir(printDupes bool) error {
	dir, _ := os.Getwd()

	fmt.Printf("Scanning: %s\n", dir);

	var appFs = afero.NewOsFs()

	return DDupe(appFs, dir, printDupes)
}

func DDupe(fsys afero.Fs, dirPrefix string, printDupes bool) error {

	result, err := ScanForDupes(fsys, dirPrefix)

	if err != nil {
		return err
	}

	fmt.Println("-------")

	if printDupes {
		var i = 0
		for i = 0; i < len(result.DupePaths); i++ {
			fmt.Println(result.DupePaths[i])
		}
		if i > 0 {
			fmt.Println("-------")
		}
	}

	fmt.Printf("Scanned %d files\n", result.FilesScanned)
	fmt.Printf("Found %d duplicates\n", len(result.DupePaths))
	fmt.Println("-------")

	if len(result.DupePaths) > 0 {
		if askYN("Delete duplicates?") {
			// do something
			fmt.Print("Deleting...\n")
			return deleteDuplicates(fsys, result.DupePaths)
		} else {
			fmt.Print("Aborted\n")
		}
	}

	return nil
}

func deleteDuplicates(fsys afero.Fs, dupes []string) error{
	for i := 0; i < len(dupes); i++ {
		fmt.Printf("Deleting: %s\n", dupes[i])
		err := fsys.Remove(dupes[i])
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func askYN(q string) bool {
	var a string = "N"
	fmt.Printf("%s (y/N): ", q)
	fmt.Scanln(&a)

	a = strings.ToLower(a)

	if a == "y" || a == "yes" {
		return true
	}

	return false
}

// ScanForDupes scans a folder recursively for duplicate files and returns a dupeScanResult
// containing all the duplicates found. The first file found is treated as the original and not
// included in the result
func ScanForDupes(fsys afero.Fs, dirPrefix string) (DupeScanResult, error) {
	result := DupeScanResult{}

	var originals = make(map[string]string)

	err := afero.Walk(fsys, dirPrefix, func(p string, info fs.FileInfo, perr error) error {

		if info.IsDir() {
			return nil
		}

		result.FilesScanned++

		f, err := fsys.Open(p)

		if err != nil {
			return err
		}

		defer f.Close()

		hash, err := HashFile(f)

		if err != nil {
			return err
		}

		if _, exists := originals[hash]; !exists {
			originals[hash] = p
		} else {
			result.DupePaths = append(result.DupePaths, p)
		}
		return nil
	})

	if err != nil {
		return DupeScanResult{}, err
	}

	return result, nil
}
