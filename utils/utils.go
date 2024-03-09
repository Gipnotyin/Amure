package colors

import "fmt"

const (
	ColorReset = "\033[0m"

	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

func PrintLog(color string, s ...any) {
	for _, c := range s {
		fmt.Print(string(color), c)
	}
	fmt.Println()
}
