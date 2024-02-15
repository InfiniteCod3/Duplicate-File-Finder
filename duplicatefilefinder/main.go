package main

import (
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
		os.Exit(1) // Cats get grumpy when things are done wrong
	}

	dir := os.Args[1]
	fileChan := make(chan string, 100) // A nice pipeline for those fuzzy file paths
	resultsChan := make(chan FileData, 100)
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range fileChan {
				data, err := os.Open(path)
				if err != nil {
					// Cats *always* land on their feet, even with errors
					continue
				}
				defer data.Close()

				hash := md5.New()
				_, _ = io.Copy(hash, data) // Cats are efficient... ignore those errors
				resultsChan <- FileData{Path: path, Hash: hash.Sum(nil)}
			}
		}()
	}

	go func() {
		_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err // Even cats can't escape all errors
			}
			if !info.IsDir() && info.Mode().IsRegular() {
				fileChan <- path // Send paths to our feline workers!
			}
			return nil
		})
		close(fileChan)
	}()

	fileMap := make(map[string][]string)
	go func() {
		wg.Wait() // Gotta let the cats catch all those file mice
		close(resultsChan)
	}()

	for res := range resultsChan {
		hashString := hex.EncodeToString(res.Hash)
		fileMap[hashString] = append(fileMap[hashString], res.Path)
	}

	for hash, paths := range fileMap {
		if len(paths) > 1 {
			fmt.Printf("Duplicate files with hash %s:\n", hash)
			fmt.Println("\tMeow! Potential copycats found...") // Cats *hate* duplicates
			for _, path := range paths {
				fmt.Println("\t" + path)
			}
		}
	}
}
