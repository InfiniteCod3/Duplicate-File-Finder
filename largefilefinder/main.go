package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

func main() {
	// Command-line flag parsing
	dirPtr := flag.String("dir", ".", "The directory to search")
	sizePtr := flag.String("size", "1MB", "Minimum file size (with unit, e.g., '1MB', '500KB', etc.) to consider 'large'")
	flag.Parse()

	// Convert size with unit to size in bytes
	sizeInBytes, err := convertSizeToBytes(*sizePtr)
	if err != nil {
		fmt.Printf("Oops! There was a problem converting the size: %v\n", err)
		return
	}

	// Validate the provided directory
	if _, err := os.Stat(*dirPtr); os.IsNotExist(err) {
		fmt.Printf("Meow! Can't find this '%s' directory. It vanished!\n", *dirPtr)
		return
	}

	// Walk the directory concurrently
	fmt.Printf("Prowling for those bulky files (larger than %s) in %s...\n", *sizePtr, *dirPtr)
	var wg sync.WaitGroup
	err = filepath.WalkDir(*dirPtr, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Hiss! Something went wrong with %s. Skipping for now.\n", path)
			return nil
		}

		if d.IsDir() {
			return nil // Skip directories
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			info, err := d.Info()
			if err != nil {
				fmt.Printf("Fur ruffled! Couldn't get info on %s. Moving on...\n", path)
				return
			}

			// Check if it's a regular file and larger than the specified size
			if info.Size() >= sizeInBytes {
				sizeStr := convertBytesToSize(info.Size())
				fmt.Printf("Aha! Caught a big one: %s (%s). Worth chasing!\n", path, sizeStr)
			}
		}()

		return nil
	})

	if err != nil {
		fmt.Printf("Hmm, something went wrong during my search: %v\n", err)
	}

	wg.Wait()
}

// convertSizeToBytes converts a size string (e.g., "1MB", "500KB") to its equivalent in bytes.
func convertSizeToBytes(sizeStr string) (int64, error) {
	sizeStr = strings.ToUpper(sizeStr)
	var unitMultiplier int64 = 1
	switch {
	case strings.HasSuffix(sizeStr, "KB"):
		unitMultiplier = 1024
		sizeStr = strings.TrimSuffix(sizeStr, "KB")
	case strings.HasSuffix(sizeStr, "MB"):
		unitMultiplier = 1024 * 1024
		sizeStr = strings.TrimSuffix(sizeStr, "MB")
	case strings.HasSuffix(sizeStr, "GB"):
		unitMultiplier = 1024 * 1024 * 1024
		sizeStr = strings.TrimSuffix(sizeStr, "GB")
	case strings.HasSuffix(sizeStr, "B"):
		sizeStr = strings.TrimSuffix(sizeStr, "B")
	}

	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parsing size failed: %w", err)
	}

	return size * unitMultiplier, nil
}

// convertBytesToSize converts a size in bytes to its equivalent string representation.
func convertBytesToSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}
