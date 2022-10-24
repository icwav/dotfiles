package main

import (
		"fmt"
		"io/ioutil"
		"os"
		"strings"
)

func main() {
//		if len(os.Args) <= 1 {
//				fmt.Printf("USAGE : %s <target_filename> \n", os.Args[0])
//				os.Exit(0)
//		}

		fileName := "testfile.txt"

		fileBytes, err := ioutil.ReadFile(fileName)

		if err != nil {
				fmt.Println(err)
				os.Exit(1)
		}

		lines := strings.Split(string(fileBytes), "\n")

		// remove the last item from the lines slice
		// which is empty
		lines = lines[:len(lines)-1]

		fmt.Println("Configuring", fileName, "...")
		for i, line := range lines {
				fmt.Println(i, line)
				if strings.Contains(line, "line1") {
						lines[i] = "L"

				}
				output := strings.Join(lines, "\n")
				err = ioutil.WriteFile(fileName, []byte (output), 0644)
				if err != nil {
					fmt.Println(err)
				}
			}
}