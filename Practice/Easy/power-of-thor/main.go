package main

import "fmt"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 * ---
 * Hint: You can use the debug stream to print initialTX and initialTY, if Thor seems not follow your orders.
 **/

func main() {
	// lightX: the X position of the light of power
	// lightY: the Y position of the light of power
	// initialTx: Thor's starting X position
	// initialTy: Thor's starting Y position
	var lightX, lightY, initialTx, initialTy int
	fmt.Scan(&lightX, &lightY, &initialTx, &initialTy)

	var moveX = lightX - initialTx
	var moveY = lightY - initialTy

	for {
		// remainingTurns: The remaining amount of turns Thor can move. Do not remove this line.
		var remainingTurns int
		fmt.Scan(&remainingTurns)

		if moveX > 0 {

			moveX -= 1

			if moveY > 0 {
				moveY -= 1
				fmt.Println("SE")
			} else if moveY < 0 {
				moveY += 1
				fmt.Println("NE")
			} else {
				fmt.Println("E")
			}
		} else if moveX < 0 {
			moveX += 1

			if moveY > 0 {
				moveY -= 1
				fmt.Println("SW")
			} else if moveY < 0 {
				moveY += 1
				fmt.Println("NW")
			} else {
				fmt.Println("W")
			}
		} else {
			if initialTy < lightY {
				fmt.Println("S")
			} else if initialTy > lightY {
				fmt.Println("N")
			}
		}
	}
}
