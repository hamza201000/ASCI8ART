package main

import "fmt"

func Split(s string) []string {
	word := []string{}
	start := 0
	for i := 0; i < len(s); i++ {
		if i!=len(s)-1 && (s[i]) == '\\'&&s[i+1]=='n' {
			fmt.Println("ok")
			word = append(word, s[start:i])
			start = i + 2
		}
	}
	if start < len(s) {
		word = append(word, s[start:])
	}
	return word
}

func main() {
	fmt.Println(len(Split("fgd\\ndgfd\\n\\n")))
	fmt.Println(len("hjh\ndf"))
}
