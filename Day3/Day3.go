package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
)

// Function that takes a slice of ints and
// returns the lowest int in the slice
func getMin(array []int) (int) {
    var min int = array[0]
    for _, value := range array {
        if min > value {
            min = value
        }
    }
    return min
}

// Function to simply return absolute value of supplied integer
func getAbs(orig int) (int) {
	if (orig < 0) {
		return -orig
	} else {
		return orig
	}
}

// Returns two string slices that contain the instructions
// w1 (wire 1) must take and w2 (wire 2) must take, respectively.
func getPuzzleInput() (w1Ops []string, w2Ops []string) {
	file, err := os.Open("puzzleinput.txt")
	if err != nil {
		log.Fatal("error")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	var w1OpsString string
	var w2OpsString string
    for scanner.Scan() {
		if (count == 0) {
			w1OpsString = scanner.Text()
		} else if (count == 1) {
			w2OpsString = scanner.Text()
		}
		count = count + 1
    }

	w1OpsSlice := strings.Split(w1OpsString, ",")
	w2OpsSlice := strings.Split(w2OpsString, ",")

	return w1OpsSlice, w2OpsSlice
}

func main() {
	w1Ops, w2Ops := getPuzzleInput()

	var w1Path [][]int
	var w2Path [][]int

	w1CurrPos := []int{0,0}
	w2CurrPos := []int{0,0}

	// Process all ops and map all coordinates for wire 1
	for i := range w1Ops {
		dir := w1Ops[i][:1]
		numCoords, err := strconv.Atoi(w1Ops[i][1:])
		if err != nil {
			log.Fatal("error")
		}
		if (dir == "U") {
			for j := 0; j < numCoords; j++ {
				w1CurrPos[1] = w1CurrPos[1] + 1
				w1Path = append(w1Path, []int{w1CurrPos[0], w1CurrPos[1]})
			}
		} else if (dir == "D") {
			for j := 0; j < numCoords; j++ {
				w1CurrPos[1] = w1CurrPos[1] - 1
				w1Path = append(w1Path, []int{w1CurrPos[0], w1CurrPos[1]})
			}
		} else if (dir == "R") {
			for j := 0; j < numCoords; j++ {
				w1CurrPos[0] = w1CurrPos[0] + 1
				w1Path = append(w1Path, []int{w1CurrPos[0], w1CurrPos[1]})
			}
		} else if (dir == "L") {
			for j := 0; j < numCoords; j++ {
				w1CurrPos[0] = w1CurrPos[0] - 1
				w1Path = append(w1Path, []int{w1CurrPos[0], w1CurrPos[1]})
			}
		}
	}

	// Process all ops and map all coordinates for wire 2
	for i := range w2Ops {
		dir := w2Ops[i][:1]
		numCoords, err := strconv.Atoi(w2Ops[i][1:])
		if err != nil {
			log.Fatal("error")
		}
		if (dir == "U") {
			for j := 0; j < numCoords; j++ {
				w2CurrPos[1] = w2CurrPos[1] + 1
				w2Path = append(w2Path, []int{w2CurrPos[0], w2CurrPos[1]})
			}
		} else if (dir == "D") {
			for j := 0; j < numCoords; j++ {
				w2CurrPos[1] = w2CurrPos[1] - 1
				w2Path = append(w2Path, []int{w2CurrPos[0], w2CurrPos[1]})
			}
		} else if (dir == "R") {
			for j := 0; j < numCoords; j++ {
				w2CurrPos[0] = w2CurrPos[0] + 1
				w2Path = append(w2Path, []int{w2CurrPos[0], w2CurrPos[1]})
			}
		} else if (dir == "L") {
			for j := 0; j < numCoords; j++ {
				w2CurrPos[0] = w2CurrPos[0] - 1
				w2Path = append(w2Path, []int{w2CurrPos[0], w2CurrPos[1]})
			}
		}
	}

	//fmt.Println(w1Path)

	var intersectionDistances[]int
	for i := range w1Path {
		for j := range w2Path {
			if (w1Path[i][0] == w2Path[j][0] && w1Path[i][1] == w2Path[j][1]) {
				intersectionDistances = append(intersectionDistances, getAbs(w1Path[i][0]) + getAbs(w1Path[i][1]))
			}
		}
	}
	fmt.Println("Part 1 Solution:")
	fmt.Println(getMin(intersectionDistances))
	
}
