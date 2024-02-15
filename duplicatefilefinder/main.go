package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"bytes"
)

// FileData: A precious little data morsel for our curious cats
type FileData struct {
	Path string
	Hash []byte // A unique pawprint for each file
}

// main: Where the kitty magic begins...
func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <directory>\n", os.Args[0])
		fmt.Println("Hiss! Use the right paw-wer next time, human!") // A gentle scolding
		os.Exit(1) 
	}

	dir := os.Args[1]

	// Channels for kitty communication
	fileChan := make(chan string, 100)    // Paths for hungry kitties
	resultsChan := make(chan FileData, 100) // Yummy morsels with hashes

	var wg sync.WaitGroup // Organizing our fluffy investigators

	// Adjusting the number of whiskers in the hunt for efficiency...
	var numWorkers = runtime.NumCPU() // More paws for more power!
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range fileChan {
				if err := processFile(path, resultsChan); err != nil {
					fmt.Fprintf(os.Stderr, "Mew! Trouble sniffing %s: %v\n", path, err)
				}
			}
		}()
	}

	// A curious kitty exploring
	go func() {
		_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err // Whiskers twitching... a troublesome path
			}
			if !info.IsDir() && info.Mode().IsRegular() { // Regular files only, please!
				fileChan <- path
			}
			return nil
		})
		close(fileChan) // No more file snacks - let's digest
	}()

	// Time to gather the kitty crew when the snacks are all eaten
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Building the scent map... each hash a unique scent marking territory
	fileMap := make(map[string][]string)
	for res := range resultsChan {
		hashString := hex.EncodeToString(res.Hash)
		fileMap[hashString] = append(fileMap[hashString], res.Path)
	}

	// Pouncing on those copycats!
	for hash, paths := range fileMap {
		if len(paths) > 1 {
			fmt.Printf("Duplicate files with hash %s:\n", hash)
			fmt.Println("\tPurr... found some potential copycats!") // Playful discovery
			for _, path := range paths {
				fmt.Println("\t" + path)
			}

			// Checking our scent marks extra carefully in case of mischief
			if !compareFiles(paths[0], paths[1]) {
				fmt.Println("Warning: Hash collision detected! Might be mischievous copycats!")
			}
		}
	}
}

// processFile: A meticulous kitty inspecting a file
func processFile(path string, resultsChan chan<- FileData) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close() // Gently close when sniffing is done

	hash := md5.New() // Our kitty's super powerful nose
	reader := bufio.NewReader(file) // For efficient nibbles of data

	if _, err := io.Copy(hash, reader); err != nil {
		return err
	}

	resultsChan <- FileData{Path: path, Hash: hash.Sum(nil)} // Sharing the findings!
	return nil
}

// compareFiles: When kitties need to double-check those scents
// They don't get fooled easily!
func compareFiles(path1, path2 string) bool {
	file1, err1 := os.Open(path1)
	if err1 != nil {
		return false // Can't open? Assume they're different
	}
	defer file1.Close()

	file2, err2 := os.Open(path2)
	if err2 != nil {
		return false
	}
	defer file2.Close()

	// Efficient comparison using buffers
	const bufferSize = 64 * 1024 // 64KB
	buffer1 := make([]byte, bufferSize)
	buffer2 := make([]byte, bufferSize)

	for {
		n1, err1 := file1.Read(buffer1)
		n2, err2 := file2.Read(buffer2)

		if err1 != err2 || !bytes.Equal(buffer1[:n1], buffer2[:n2]) {
			return false // Mismatch found
		}

		if err1 == io.EOF && err2 == io.EOF {
			return true // Files are identical
		}
	}
} 
