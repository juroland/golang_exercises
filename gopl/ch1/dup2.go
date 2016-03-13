package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FilesCount struct {
	filenames []string
	count     int
}

func main() {
	counts := make(map[string]*FilesCount)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, filesCount := range counts {
		if filesCount.count > 1 {
			fmt.Printf("%d\t%s\n", line, filesCount.count)
			fmt.Printf(strings.Join(filesCount.filenames, "\n"))
			fmt.Println()
		}
	}
}

func countLines(f *os.File, counts map[string]*FilesCount) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		_, ok := counts[line]
		if !ok { 
			counts[line] = new(FilesCount)
		}
		counts[line].count++
		counts[line].filenames = append(counts[line].filenames, f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}
