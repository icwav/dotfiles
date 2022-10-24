package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	readFile, err := os.OpenFile("testfile.txt",os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	
	data := ("hsfdsdf")

	for fileScanner.Scan() {
		file = fileScanner.Text()
		if (file == "hello1") {
			_, err = readFile.WriteString(data)
			if err != nil{
				fmt.Println(err)
			}
		}
	readFile.Close()
	}
}