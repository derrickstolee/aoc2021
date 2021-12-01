package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var increases int = 0
	var last_depth int = -1

	line, err := reader.ReadString('\n')

	for err == nil {
		line = strings.TrimSpace(line)
		number, nerr := strconv.Atoi(line)

		if last_depth > 0 && last_depth < number {
			increases++
		}

		if nerr != nil {
			fmt.Printf("Line: %s, nerr: %s", line, nerr);
			return
		}

		line, err = reader.ReadString('\n')
		last_depth = number
	}

	fmt.Printf("Increases: %d\n", increases)
}
