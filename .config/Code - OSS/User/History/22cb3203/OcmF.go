package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	readFile, err := os.Open("testfile.txt",os.O_APPEND|os.O_WRONGLY, 0777)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	
	data := []byte("hsfdsdf")

	for fileScanner.Scan() {
		if (fileScanner.Text() == "hello" ) {
			_, err = fmt.Fprintln(readFile , data)
			if err != nil{
				fmt.Println(err)
			}
		}
	readFile.Close()
	}
}