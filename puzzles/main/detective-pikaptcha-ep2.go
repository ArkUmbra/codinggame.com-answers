package main

import (
	"strconv"
	"fmt"
)


const EMPTY = "0"
const WALL = "#"


/**
	Solution to puzzle at: https://www.codingame.com/ide/puzzle/detective-pikaptcha-ep1
 */
func main() {
	var width, height int
	fmt.Scan(&width, &height)

	// 2d slices
	mapping := make([][]string, 0)

	// read inputs
	for i := 0; i < height; i++ {
		var line string
		fmt.Scan(&line)

		j := 0
		// slice of row
		row := make([]string, 0)
		for _, char := range line {
			row = append(row, string(char))
			j = j + 1
		}
		mapping = append(mapping, row)
	}

	// process map
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var charToPrint string

			if mapping[y][x] == EMPTY {
				numPassable := findPassableAdjacent(mapping, x, y, width, height)
				//charToPrint = EMPTY
				charToPrint = strconv.Itoa(numPassable)
			} else {
				charToPrint = WALL
			}


			fmt.Printf(charToPrint)
		}

		fmt.Printf("\n")
	}
}


func findPassableAdjacent(mapping [][]string, x int, y int, maxW int, maxH int) int {
	totalPassable := 0

	// check up
	if isPassable(mapping, x, y-1, maxW, maxH) {
		totalPassable = totalPassable + 1
	}

	// check right
	if isPassable(mapping, x + 1, y, maxW, maxH) {
		totalPassable = totalPassable + 1
	}

	// check left
	if isPassable(mapping, x -1, y, maxW, maxH) {
		totalPassable = totalPassable + 1
	}

	// check down
	if isPassable(mapping, x, y + 1, maxW, maxH) {
		totalPassable = totalPassable + 1
	}

	return totalPassable
}

func isPassable(mapping [][]string, x int, y int, maxW int, maxH int) bool {
	if x < 0 || x > maxW -1 {
		return false
	}

	if y < 0 || y > maxH -1 {
		return false
	}

	return mapping[y][x] == EMPTY
}