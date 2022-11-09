package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	}else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprint(os.Stderr, "dub2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	
	for word, n := range counts {
		if n > 1 {
			fmt.Printf("Quantidade: %d, Palavra: %s\n", n, word)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f);
	
	for input.Scan() {
		counts[input.Text() + "\tArquivo: " +f.Name()]++
	}
}