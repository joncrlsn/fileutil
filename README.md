fileutil
========

Golang text file reading and writing utilities I've written based on examples provided in various public sources on the web.

/*
 * Reads a file line by line into a channel
 *
 * c := fileutil.ReadLinesChannel(fileName)
 * for line := range c {
 *   fmt.Printf("  Line: %s\n", line)
 * }
 */
func ReadLinesChannel(fileName string) <-chan string

/*
 * Reads a file line by line into an array
 *
 *  lines, err := fileutil.ReadLinesArray(fileName)
 *  if err != nil {
 *      log.Fatalf("readLines: %s\n", err)
 *  }
 *  for i, line := range lines {
 *      fmt.Printf("  Line: %d %s\n", i, line)
 *  }
 */
func ReadLinesArray(fileName) ([]string, err)

// Writes the lines to the given file.
func WriteLinesArray(lines []string, path string) error {

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

// Generates a temporary file path for use in testing or whatever
func TempFileName(prefix, suffix string) string {
    randBytes := make([]byte, 16)
    rand.Read(randBytes)
    return filepath.Join(os.TempDir(), prefix + hex.EncodeToString(randBytes) + suffix)
}
