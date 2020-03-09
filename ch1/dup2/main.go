package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintln(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Println(line)
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines (f *os.File, counts map[string]int) {
	//fmt.Print(f);
	input :=bufio.NewScanner(f)
	for input.Scan() {
		fmt.Printf("%d\t%s\n", input.Bytes(),input.Text());
		counts[input.Text()]++
	}
}
