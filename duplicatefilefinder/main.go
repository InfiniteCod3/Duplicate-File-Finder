package main

import (
	"bufio" // For efficient file sniffing
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

type FileData struct {
	Path string
	Hash []byte
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <directory>\n", os.Args[0])
		os.Exit(1) // Gotta hiss at incorrect usage!
	}

	dir := os.Args[1]

	// Prepare a cozy pipeline for file paths and results
	fileChan := make(chan string, 100) 
	resultsChan := make(chan FileData, 100)

	var wg sync.WaitGroup // Manage those busy worker cats

	// Unleash a few skilled worker cats for file inspection
	const numWorkers = 5
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range fileChan {
				// Handle errors with feline finesse
				if err := processFile(path, resultsChan); err != nil {
					fmt.Fprintf(os.Stderr, "Hiss! Couldn't process %s: %v\n", path, err)
				}
			}
		}()
	}

	// Time for a stroll to find interesting files
	go func() {
		_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err // Whiskers twitch... something's wrong 
			}
			if !info.IsDir() && info.Mode().IsRegular() {
				fileChan <- path // Send a 'toy' to the worker cats
			}
			return nil
		})
		close(fileChan) // No more toys after this!
	}()

	// Gather results gracefully... cats always finish their task
	go func() {
		wg.Wait() 
		close(resultsChan) 
	}()

	// Build a scent map to track those copycats
	fileMap := make(map[string][]string)
	for res := range resultsChan {
		hashString := hex.EncodeToString(res.Hash)
		fileMap[hashString] = append(fileMap[hashString], res.Path)
	}

	// Time to pounce on duplicates!
	for hash, paths := range fileMap {
		if len(paths) > 1 {
			fmt.Printf("Duplicate files with hash %s:\n", hash)
			fmt.Println("\tPurr... found some potential copycats!") 
			for _, path := range paths {
				fmt.Println("\t" + path)
			}
		}
	}
}

// processFile: Where a cat carefully examines a file's 'scent'
func processFile(path string, resultsChan chan<- FileData) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	hash := md5.New()
	// Time to sniff that file with our expert noses
	reader := bufio.NewReader(file) 
	if _, err := io.Copy(hash, reader); err != nil {
		return err
	}

	resultsChan <- FileData{Path: path, Hash: hash.Sum(nil)}
	return nil
}
