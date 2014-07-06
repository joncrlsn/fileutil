package fileutil

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "encoding/hex"
    "path/filepath"
    "crypto/rand"
)

/*
 * Reads a text file line by line into a channel.
 *
 * c := fileutil.ReadLinesChannel(fileName)
 * for line := range c {
 *   fmt.Printf("  Line: %s\n", line)
 * }
 */
func ReadLinesChannel(filePath string) <-chan string {
    c := make(chan string)
    file, err := os.Open(filePath)
    if err != nil {
        log.Panic(err)
    }
    go func() {
        defer file.Close()
        scanner := bufio.NewScanner(file)
        for scanner.Scan(){
            c <- scanner.Text()
        }
        close(c)
    }()
    return c
}


/*
 * Reads a text file line by line into an array.  Not recommended
 * for use with large files.
 *
 *  lines, err := fileutil.ReadLinesArray(filePath)
 *  if err != nil {
 *      log.Fatalf("readLines: %s\n", err)
 *  }
 *  for i, line := range lines {
 *      fmt.Printf("  Line: %d %s\n", i, line)
 *  }
 */
func ReadLinesArray(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

/*
 * Writes the lines to the given file.
 */
func WriteLinesArray(lines []string, path string) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()

    w := bufio.NewWriter(file)
    for _, line := range lines {
        fmt.Fprintln(w, line)
    }
    return w.Flush()
}

// Returns whether or not the given file or directory exists
func Exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}

// Generates a temporary filename for use in testing or whatever
func TempFileName(prefix, suffix string) string {
    randBytes := make([]byte, 16)
    rand.Read(randBytes)
    return filepath.Join(os.TempDir(), prefix + hex.EncodeToString(randBytes) + suffix)
}
