package main


import (
    "bufio"
    "fmt"
    "os"
)

// print out result
func printResult(counts map[string]int) {
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
     }
}

func countLines(f *os.File, counts map[string]int)  {
    input := bufio.NewScanner(f)
    sep := " # "
    for input.Scan() {
        counts[f.Name() + sep + input.Text()]++
    }
}

// dup4 prints names of files in which duplicated lines occur
func main() {
    counts := make(map[string]int)
    files := os.Args[1:]

    if len(files) == 0 {
        countLines(os.Stdin, counts)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup4: %v\n", err)
                continue
            }
            countLines(f, counts)
            f.Close()
        }
    }

    printResult(counts)
}
