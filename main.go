package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("error :", err)
	}
	sacnneer := bufio.NewScanner(file)

	asci := [][]string{}
	asci_line := []string{}
	//str:=""
	b := false
	for sacnneer.Scan() {
		if sacnneer.Text() == "" && b {
			asci = append(asci, asci_line)
			asci_line = []string{}
			continue
		}
		b = true
		asci_line = append(asci_line, sacnneer.Text())
	}
	str := "hamza"
	//fmt.Println(asci[0])

	m := 0

	for i := 0; i < 8; i++ {
		for j := 0; j < len(str); j++ {
			m = int(str[j] - 32)
			fmt.Print(asci[m][i])
		}
		fmt.Println()
	}

}
