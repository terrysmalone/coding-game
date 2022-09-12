package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// N: Number of elements which make up the association table.
	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	// Q: Number Q of file names to be analyzed.
	var Q int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &Q)

	typeMap := make(map[string]string)

	for i := 0; i < N; i++ {
		// EXT: file extension
		// MT: MIME type.
		var EXT, MT string
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &EXT, &MT)

		typeMap[strings.ToLower(EXT)] = MT
	}
	for i := 0; i < Q; i++ {
		scanner.Scan()
		FNAME := scanner.Text()

		extensionWithDot := filepath.Ext(FNAME)
		_, i := utf8.DecodeRuneInString(extensionWithDot)
		extension := strings.ToLower(extensionWithDot[i:])

		mimeType := typeMap[extension]

		if mimeType != "" {
			fmt.Println(mimeType)
		} else {
			fmt.Println("UNKNOWN")
		}

	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")

	// For each of the Q filenames, display on a line the corresponding MIME type. If there is no corresponding type, then display UNKNOWN.
	//fmt.Println("UNKNOWN")
}
