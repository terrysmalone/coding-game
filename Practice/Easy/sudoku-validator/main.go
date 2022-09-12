package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var inputs []string

	horizontal := make([][]int64, 9)
	vertical := make([][]int64, 9)
	quadrant := make([][]int64, 9)

	for column := 0; column < 9; column++ {
		scanner.Scan()
		inputs = strings.Split(scanner.Text(), " ")
		for row := 0; row < 9; row++ {
			num, _ := strconv.ParseInt(inputs[row], 10, 32)

			horizontal[column] = append(horizontal[column], num)
			vertical[row] = append(vertical[row], num)

			if column < 3 {
				if row < 3 {
					quadrant[0] = append(quadrant[0], num)
				} else if row < 6 {
					quadrant[1] = append(quadrant[1], num)
				} else {
					quadrant[2] = append(quadrant[2], num)
				}
			} else if column < 6 {
				if row < 3 {
					quadrant[3] = append(quadrant[3], num)
				} else if row < 6 {
					quadrant[4] = append(quadrant[4], num)
				} else {
					quadrant[5] = append(quadrant[5], num)
				}
			} else {
				if row < 3 {
					quadrant[6] = append(quadrant[6], num)
				} else if row < 6 {
					quadrant[7] = append(quadrant[7], num)
				} else {
					quadrant[8] = append(quadrant[8], num)
				}
			}
		}
	}

	isValid := true

	for i := 0; i < 9; i++ {
		if duplicateInArray(horizontal[i]) {
			isValid = false
			break
		}

		if duplicateInArray(vertical[i]) {
			isValid = false
			break
		}

		if duplicateInArray(quadrant[i]) {
			isValid = false
			break
		}
	}

	if isValid {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func duplicateInArray(arr []int64) bool {
	visited := make(map[int64]bool, 0)
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] == true {
			return true
		} else {
			visited[arr[i]] = true
		}
	}
	return false
}
