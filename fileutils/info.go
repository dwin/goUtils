package fileutils

import (
	"errors"
	"os"
	"path/filepath"
)

var (
	// ErrNotDirectory indicates the given path does not lead to a directory, path must not be file.
	ErrNotDirectory = errors.New("PathError must give path to directory not file")
)

// DirSize takes a directory path, then returns the total size of all files and error
func DirSize(path string) (size int64, err error) {
	err = filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

// DirFilenames takes a directory path, returns a list of filenames and error
func DirFilenames(path string) (info []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	// If Not Directory return error
	if !fi.IsDir() {
		return nil, ErrNotDirectory
	}
	defer f.Close()
	return f.Readdirnames(0)
}

/*
func DirFileHashes(path string, hashType string, threads int) (info []map[string]string, err error) {
	// Determine Hashing Method
	var method string
	switch hashType {
	case "Blake2b", "blake2b", "b2b":
		method = "Blake2b"
	default:
		method = "Blake2b"
	}

	// Get Directory Filenames
	filenames, err := DirFilenames(path)
	if err != nil {
		return
	}

	// Create appropriate sized channel for hashing
	jobs := make(chan string, len(filenames))
	results := make(chan map[string]string, len(filenames))

	// Start hashers
	for thread := 0; thread < threads; thread++ {
		go hashFile(thread, method, jobs, results)
	}
	var jobcount int
	// Loop through all filenames returned from getDirFiles and send to jobs queue
	for i := 0; i < len(filenames); i++ {
		// Create file and send to fileHash
		//fmt.Println("Sending file: ", i)
		jobs <- path + "/" + filenames[i]
		jobcount++
	}
	close(jobs)
	for i := 0; i < len(filenames)-1; i++ {
		f := <-results
		//fmt.Println("received result ", i)
		files = append(files, f)
	}
	return
}

func hashFile(worker int, method string, jobs <-chan string, results chan<- map[string]string) {
	for j := range jobs {
		if method == "Blake2b" {
			result := make(map[string]string)
			hash := blake2b.New512()
			file, err := os.Open(j)
			if err != nil {
				fmt.Println("Cannot open file, err: ", err)

				continue
			}
			defer file.Close()
			_, err = io.Copy(hash, file)
			if err != nil {
				fmt.Println("Cannot hash file")

				continue
			}
			resultBlake2bHEX := hex.EncodeToString(hash.Sum(nil))
			result[j] = resultBlake2bHEX
			results <- result
		}

	}

}
*/
