package asciart

import "fmt"

func PrintAsci(str []string, asci_table [][]string) {
	for k := 0; k < len(str); k++ {
		for i := 0; i < 8; i++ {
			for j := 0; j < len(str[k]); j++ {
				m := (str[k][j] - 32)
				fmt.Print(asci_table[m][i])
			}
			fmt.Println()
			if len(str[k]) == 0 {
				break
			}
		}
	}
}
