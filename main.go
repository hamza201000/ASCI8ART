package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	str := os.Args[1]
	//ne := strings.ReplaceAll(str, `\n`, "\n")
	newstring := strings.Split(str, "\\n")
	//fmt.Println(str)
	
	// fmt.Println(asci[0])
	/*b1 := [][]string{}
	newlne := []string{"\n"}
	for i := 0; i < len(str); i++ {
		m := (str[i] - 32)
		if str[i] == '\n' {
			b1 = append(b1, newlne)
			continue
		}
		b1 = append(b1, asci[m])
	}
	fmt.Println(b1[0][0])*/
	//s := 0
	// m:=0
	//fmt.Println(len(newstring))
	if newstring[len(newstring)-1]=="" {
		newstring=newstring[:len(newstring)-1]
	}
	for k := 0; k < len(newstring); k++ {
		for i := 0; i < 8; i++ {
			for j := 0; j < len(newstring[k]); j++ {
				m := (newstring[k][j] - 32)
				fmt.Print(asci[m][i])
			}
			fmt.Println()
			if len(newstring[k])==0{
				break
			}
		}
		 
	}
}
