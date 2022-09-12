package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type Direction int64

const (
	Up Direction = iota
	Down
	Left
	Right
	None
)

type arrow struct {
	direction Direction
	xPos      int
	yPos      int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var height, width int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &height, &width)

	var arrows []arrow

	for i := 0; i < height; i++ {
		scanner.Scan()
		line := scanner.Text()

		letters := []rune(line)

		for x, letter := range letters {
			if letter != '.' {
				arrows = append(arrows, arrow{
					direction: getDirection(letter),
					xPos:      x,
					yPos:      i,
				})
			}
		}
	}

	moveCount := 0

	for {
		// Move arrows
		for i, _ := range arrows {
			if arrows[i].direction == Up {
				arrows[i].yPos--
				if arrows[i].yPos < 0 {
					arrows[i].yPos = height - 1
				}
			} else if arrows[i].direction == Down {
				arrows[i].yPos++
				if arrows[i].yPos > height-1 {
					arrows[i].yPos = 0
				}
			} else if arrows[i].direction == Left {
				arrows[i].xPos--
				if arrows[i].xPos < 0 {
					arrows[i].xPos = width - 1
				}
			} else if arrows[i].direction == Right {
				arrows[i].xPos++
				if arrows[i].xPos > width-1 {
					arrows[i].xPos = 0
				}
			}
		}

		moveCount++

		var toDestroy []int

		// check for clashes
		for i := 0; i < len(arrows)-1; i++ {
			for j := i + 1; j < len(arrows); j++ {
				if arrows[i].xPos == arrows[j].xPos && arrows[i].yPos == arrows[j].yPos {
					toDestroy = append(toDestroy, i)
					toDestroy = append(toDestroy, j)
				}
			}
		}
		toDestroy = removeDuplicates(toDestroy)

		// Sort by descending
		sort.Slice(toDestroy, func(i, j int) bool {
			return toDestroy[i] > toDestroy[j]
		})

		// Remove them
		for i := 0; i < len(toDestroy); i++ {
			arrows[toDestroy[i]] = arrows[len(arrows)-1]
			arrows = arrows[:len(arrows)-1]
		}

		if len(arrows) <= 0 {
			fmt.Println(moveCount)
			break
		}
	}
}

func getDirection(letter rune) Direction {
	switch letter {
	case 'v':
		return Down
	case '>':
		return Right
	case '<':
		return Left
	case '^':
		return Up
	}

	return None
}

func removeDuplicates(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
