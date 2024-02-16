package main

import (
    "bufio"
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "runtime"
    "sync"
)

// FileData: A tasty morsel of file information 
type FileData struct {
    Path string
    Hash []byte
}

func main() {
    if len(os.Args) != 2 {
        fmt.Printf("Meow! Usage: %s <directory>\n", os.Args[0]) 
        os.Exit(1) // Hiss! Wrong way to call this kitty program
    }

    dir := os.Args[1]

    fileChan := make(chan string, 100) // A magical yarn ball to chase files
    resultsChan := make(chan FileData, 100) 

    var wg sync.WaitGroup  // Keeping track of all my busy paws
    var numWorkers = runtime.NumCPU() 
    for i := 0; i < numWorkers; i++ { 
        wg.Add(1)
        go func() {
            defer wg.Done()
            for path := range fileChan { 
                if err := processFile(path, resultsChan); err != nil {
                    fmt.Fprintf(os.Stderr, "Furball! Can't play with %s: %v\n", path, err)
                }
            }
        }() 
    }

    // **Exploring territory!** 
    go func() {
        _ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error { 
            if err != nil {
                return err  // Uh-oh, stumbled on a prickly bush!
            }
            if !info.IsDir() && info.Mode().IsRegular() {
                fileChan <- path  // Found a potential toy!
            }
            return nil 
        })
        close(fileChan) 
    }()

    // **Waiting for the hunt to be over** 
    go func() {
        wg.Wait()
        close(resultsChan) 
    }()

    // **Sorting through the treasures**
    fileMap := make(map[string][]string)
    for res := range resultsChan {
        hashString := hex.EncodeToString(res.Hash) // Turning scent into readable marks
        fileMap[hashString] = append(fileMap[hashString], res.Path) 
    }

    // **Did I find any doubles?**
    for hash, paths := range fileMap { 
        if len(paths) > 1 {
            fmt.Printf("Paws up! Duplicate scents with hash %s:\n", hash) 
            for _, path := range paths {
                fmt.Println("\t" + path) 
            }

            if !compareFiles(paths[0], paths[1]) { // Extra cautious kitty check
                fmt.Println("Warning: My whiskers are tingling! Possible mix-up!")
            }
        }
    }
}

// processFile: The hunter who sniffs out file secrets
func processFile(path string, resultsChan chan<- FileData) error {
   file, err := os.Open(path)
   if err != nil {
      return err // Oops, couldn't get my paws on that file!
   }
   defer file.Close()

   hash := md5.New() // My super-sensitive nose for catnip... I mean, data 
   reader := bufio.NewReader(file)

   if _, err := io.Copy(hash, reader); err != nil {
       return err // Hmmm, this doesn't smell right...
   }

   resultsChan <- FileData{Path: path, Hash: hash.Sum(nil)} // Sharing the catch!
   return nil
}

// compareFiles: Checking if my toys are *really* the same
func compareFiles(path1, path2 string) bool {
    file1, err1 := os.Open(path1)
    if err1 != nil {
        return false // Can't open up those toys for comparison!
    }
    defer file1.Close()

    file2, err2 := os.Open(path2)
    if err2 != nil {
        return false 
    }
    defer file2.Close()

    const bufferSize = 64 * 1024  // Playtime has to be efficient!
    buffer1 := make([]byte, bufferSize)
    buffer2 := make([]byte, bufferSize)

    for {
        n1, err1 := file1.Read(buffer1)
        n2, err2 := file2.Read(buffer2)

        if err1 != err2 || n1 != n2 || !bytes.Equal(buffer1[:n1], buffer2[:n2]) {
            return false  // Hmm, these feel different...
        }

        if err1 == io.EOF { 
            return true // They match! Time for treats!
        }
    }
}
