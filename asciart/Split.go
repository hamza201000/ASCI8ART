package asciart

import (
	"strings"
)

func Split_with_new_line(str string) []string {
	word := []string{}
	temp_str := ""
	str = strings.ReplaceAll(str, `\n`, "\n")
	for i := 0; i < len(str); i++ {
		if str[i] == '\n' {
			word = append(word, temp_str)
			temp_str = ""
		} else {
			temp_str += string(str[i])
		}
	}
	if len(temp_str) > 0 {
		word = append(word, temp_str)
	}
	return word
}
