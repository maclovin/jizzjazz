package shared

import "fmt"

const (
	BoldFont  = "\033[1m"
	ResetFont = "\033[0m"
)

func PrintTable(headers []string, rows [][]string) {
	fmt.Println()

	for _, header := range headers {
		fmt.Printf("%s%-10s%s", BoldFont, header, ResetFont)
	}

	fmt.Println()

	for _, row := range rows {
		for _, col := range row {
			fmt.Printf("%-10s", col)
		}

		fmt.Println()
	}
}
