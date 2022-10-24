package main

import (
		"fmt"
		"io/ioutil"
		"os"
		"strings"
)

func main() {

		fileName := os.Args[1]

		fileBytes, err := ioutil.ReadFile("testfile.txt")

		if err != nil {
				fmt.Println(err)
				os.Exit(1)
		}

		lines := strings.Split(string(fileBytes), "\n")

		// remove the last item from the lines slice
		// which is empty
		lines = lines[:len(lines)-1]

		// len() function will count the total number of lines
		fmt.Println(fileName, "has a total of", len(lines), "lines")
		for i, line := range lines {
				// i = i + 1 // uncomment to start from 1 instead of 0
				fmt.Println(i, line)
		}

}