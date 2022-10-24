package main

import (
/*	"io/ioutil"
	"log"
	"net/http"
	"fmt"
	"os"
*/
)




func main() {
	new, err := os.Create("telnet")
	if err != nil {
		log.Fatal(err)
	}

	trequest, err := http.Get("https://raw.githubusercontent.com/CSUSB-CISO/godaedalus/main/Configurations/telnet")
	if err != nil {
		log.Fatalln(err)
	}

	tbody, err := ioutil.ReadAll(trequest.Body)
	if err != nil {
		log.Fatalln(err)
	}

	file2write := string(tbody)
	

	os.Open(newfile)
	if err != nil {
		log.Fatal(err)
	}
*/
}