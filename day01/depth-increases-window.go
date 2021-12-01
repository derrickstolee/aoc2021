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
	var last_sum int = -1
	var depths [4] int
	var cur_count int = 0

	const window = 3

	line, err := reader.ReadString('\n')

	for err == nil {
		line = strings.TrimSpace(line)
		number, nerr := strconv.Atoi(line)

		if nerr != nil {
			fmt.Printf("Line: %s, nerr: %s", line, nerr);
			return
		}

		depths[cur_count] = number
		cur_count++

		if cur_count == window {
			last_sum = 0

			for i := 0; i < window; i++ {
				last_sum += depths[i]
			}
		}

		if cur_count == window + 1 {
			next_sum := 0
			for i := 1; i <= window; i++ {
				next_sum += depths[i]
			}

			if last_sum < next_sum {
				increases++
			}

			last_sum = next_sum

			for i := 0; i < window; i++ {
				depths[i] = depths[i + 1]
			}
			cur_count--
		}

		line, err = reader.ReadString('\n')
	}

	fmt.Printf("Increases: %d\n", increases)
}
