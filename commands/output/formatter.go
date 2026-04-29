package output

import (
	"fmt"
	"strings"
)

// PrintTable prints a table to stdout with the given header and data.
func PrintTable(header []string, data [][]string) {
	if len(header) == 0 {
		return
	}

	widths := make([]int, len(header))
	for i, h := range header {
		widths[i] = len(h)
	}

	for _, row := range data {
		for i, val := range row {
			if i < len(widths) && len(val) > widths[i] {
				widths[i] = len(val)
			}
		}
	}

	format := ""
	for _, w := range widths {
		format += fmt.Sprintf("%%-%ds   ", w)
	}
	format = strings.TrimSpace(format) + "\n"

	// Print Header
	headerArgs := make([]interface{}, len(header))
	for i, h := range header {
		headerArgs[i] = strings.ToUpper(h)
	}
	fmt.Printf(format, headerArgs...)

	// Print Separator (optional, but makes it cleaner)
	// for _, w := range widths {
	// 	fmt.Print(strings.Repeat("-", w), "   ")
	// }
	// fmt.Println()

	// Print Rows
	for _, row := range data {
		rowArgs := make([]interface{}, len(row))
		for i, val := range row {
			rowArgs[i] = val
		}
		fmt.Printf(format, rowArgs...)
	}
}

// PrintVertical prints key-value pairs vertically.
func PrintVertical(data [][]string) {
	if len(data) == 0 {
		return
	}

	maxKeyWidth := 0
	for _, row := range data {
		if len(row) > 0 && len(row[0]) > maxKeyWidth {
			maxKeyWidth = len(row[0])
		}
	}

	format := fmt.Sprintf("%%-%ds : %%s\n", maxKeyWidth)
	for _, row := range data {
		if len(row) >= 2 {
			fmt.Printf(format, row[0], row[1])
		} else if len(row) == 1 {
			fmt.Printf(format, row[0], "")
		}
	}
}
