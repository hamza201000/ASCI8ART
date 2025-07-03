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
		if sacnneer.Text() == "" && !b {
			continue

		}
		if sacnneer.Text() == "" {
			asci = append(asci, asci_line)
			asci_line = []string{}

		} else {

			asci_line = append(asci_line, sacnneer.Text())
			b = true
		}

	}

	str := "Hello\nThere"
	//fmt.Println(asci[0])
	b1 := [][]string{}
	newlne := []string{"\n"}
	for i := 0; i < len(str); i++ {
		m := (str[i] - 32)
		if str[i] == '\n' {
			b1 = append(b1, newlne)
			continue
		}
		b1 = append(b1, asci[m])
	}
	fmt.Println(b1)
	for i := 0; i < 8; i++ {
		for j := 0; j < len(b1); j++ {
			
			//m := (str[j] - 32)
			fmt.Print(b1[j][i])
		}
		fmt.Println()
	}

}
