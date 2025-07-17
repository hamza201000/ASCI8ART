package main

import (
	"bufio"
	"fmt"
	"os"
	"asciart/asciart"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : go run main.go \"Your text here\"")
		return
	}
	file, err := os.Open("standard.txt")
	if err != nil {
		file, err = os.Open("shadow.txt")
		if err != nil {
			file, err = os.Open("thinkertoy.txt")
			if err != nil {
				fmt.Println("Error: failed to open any files")
				return
			}
		}
	}
	defer file.Close()
	Scanner := bufio.NewScanner(file)
	asci_table := asciart.ParseAsci(Scanner)
	newstring := asciart.Split_with_new_line(os.Args[1])
	asciart.PrintAsci(newstring, asci_table)
}
