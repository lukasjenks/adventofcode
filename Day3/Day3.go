package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
)

func getPuzzleInput() (w1Ops []string, w2Ops []string) {
	file, err := os.Open("puzzleinput.txt")
	if err != nil {
		log.Fatal("error")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	w1OpsString := ""
	w2OpsString := ""
    for scanner.Scan() {
		if (count == 0) {
			w1OpsString = scanner.Text()
		} else if (count == 1) {
			w2OpsString = scanner.Text()
		}
		count = count + 1
    }

	w1OpsArr := strings.Split(w1OpsString, ",")
	w2OpsArr := strings.Split(w2OpsString, ",")

	return w1OpsArr, w2OpsArr
}

func main() {
	w1Ops, w2Ops := getPuzzleInput()
	fmt.Println(w1Ops, w2Ops)
}
