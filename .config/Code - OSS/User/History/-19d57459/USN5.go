package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
	"os"
)

func main() {
	trequest, err := http.Get("https://raw.githubusercontent.com/CSUSB-CISO/godaedalus/main/Configurations/telnet")
	if err != nil {
		log.Fatalln(err)
	}

	tbody, err := ioutil.ReadAll(trequest.Body)
	if err != nil {
		log.Fatalln(err)
	}

string(tbody)
	
	newfile, err := os.Create("telnet")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(newfile)

/*	os.Open(newfile)
	if err != nil {
		log.Fatal(err)
	}
*/
}