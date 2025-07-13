// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	file, err := os.Open("standard.txt")
// 	if err != nil {
// 		fmt.Println("error :", err)
// 	}
// 	sacnneer := bufio.NewScanner(file)

// 	asci := []string{}
// 	// asci_line := ""
// 	// str:=""
	
// 	for sacnneer.Scan() {
// 			asci = append(asci, sacnneer.Text())
// 	}
// 	// fmt.Println(asci[4])
// 	str := "fsduh\nghdfg"
// 	ne := strings.ReplaceAll(str, `\n`, "\n")
// 	newstring := strings.Split(ne, "\n")

// 	fmt.Println(len(asci))
// 	// fmt.Println(asci[0])
// 	/*b1 := [][]string{}
// 	newlne := []string{"\n"}
// 	for i := 0; i < len(str); i++ {
// 		m := (str[i] - 32)
// 		if str[i] == '\n' {
// 			b1 = append(b1, newlne)
// 			continue
// 		}
// 		b1 = append(b1, asci[m])
// 	}
// 	// fmt.Println(b1[0][0])

// 	// m:=0*/
// 	// s := 0
// 	res := ""
// 	for _, char := range newstring {
// 	    for i := 0; i < 8; i++ {
// 			for j := 0; j < len(char); j++ {
// 				g := i + (int(char[j]-32) * 9)
// 				res += asci[g]
// 			}
// 			res += "\n"
// 		}
// 		// fmt.Println()
// 	}
// 	fmt.Println(res)
// }
