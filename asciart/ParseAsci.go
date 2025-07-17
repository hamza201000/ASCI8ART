package asciart

import "bufio"

func ParseAsci(Scanner *bufio.Scanner) [][]string {
	asci := [][]string{}
	asci_line := []string{}
	b := false
	for Scanner.Scan() {
		if Scanner.Text() == "" && !b {
			continue
		}
		if Scanner.Text() == "" {
			asci = append(asci, asci_line)
			asci_line = []string{}
		} else {
			asci_line = append(asci_line, Scanner.Text())
			b = true
		}
	}
	return asci
}
