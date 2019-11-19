package main

import "fmt"
import "os"
import "bufio"
import "strings"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type Letter struct {
	init bool
	lines []string
}

const ELEMENTS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ?"
const NUMBER_OF_ELEMENTS = 26

var indexToLetter map[int]string
var asciiMap map[string]Letter


func main() {
	indexToLetter:= make(map[int]string)
	asciiMap := make(map[string]Letter)


	// Read Inputs
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// input row 1 - letter width
	var L int
	scanner.Scan()
	fmt.Sscan(scanner.Text(),&L)

	// input row 2 - letter height
	var H int
	scanner.Scan()
	fmt.Sscan(scanner.Text(),&H)

	// input row 3 - target output text
	scanner.Scan()
	targetText := scanner.Text()



	// Prep letter map
	// asciiMap := make(map)string]Letter
	for index, char := range ELEMENTS {
		elem := string(char)
		//fmt.Println(index, elem)
		indexToLetter[index] = elem

		//var lines [H]string
		linesSlice := make([]string, H)
		//fmt.Println(len(linesSlice))
		asciiMap[elem] = Letter{true, linesSlice}
	}


	// input remaining rows - ascii alphabet
	for y := 0; y < H; y++ {
		scanner.Scan()
		row := scanner.Text()

		// split row into character elements
		for x := 0; x <= NUMBER_OF_ELEMENTS; x++ {
			startPos := x * L
			endPos := startPos + L

			cut := row[startPos:endPos]
			//fmt.Println(cut)
			thisLetter := indexToLetter[x]
			asciiMap[thisLetter].lines[y] = cut
		}

	}


	// iterate the target text and print
	for row := 0; row < H; row++ {
		for _, char := range targetText {
			character := string(char)
			asciiRow := getLetter(character, row, asciiMap)
			fmt.Printf(asciiRow)
		}
		fmt.Printf("\n")
	}
}

func getLetter(letter string, height int, asciiMapping map[string]Letter) string {
	upper := strings.ToUpper(letter)
	ascii := asciiMapping[upper]

	// If the input letter is not recognised, should show question mark
	if ascii.init == false {
		ascii = asciiMapping["?"]
	}

	return ascii.lines[height]
}