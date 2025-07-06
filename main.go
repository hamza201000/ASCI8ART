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
	// str:=""
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
	// fmt.Println(asci[0])
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
	fmt.Println(b1[0][0])
	s:=0
	for i := 0; i < 8; i++ {
		for j := s; j < len(b1); j++ {
			if b1[j][0] == "\n" {
				continue
			}
			// m := (str[j] - 32)
			fmt.Print(b1[j][i])
			if j != len(b1)-1 && i == 7 && b1[j+1][0] == "\n" {
				s=j+1
				i = 0
				break
			}
			if j != len(b1)-1 && b1[j+1][0] == "\n" {
				j = 0
				break
			}
			
		}
		fmt.Println()
	}
}
