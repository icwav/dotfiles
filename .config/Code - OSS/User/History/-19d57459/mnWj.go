package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)

func main() {
	trequest, err := http.Get("https://raw.githubusercontent.com/CSUSB-CISO/godaedalus/main/Configurations/telnet?token=GHSAT0AAAAAABZQGHPBLQZAXJDZ3N2FCTSSY2DR3ZQ")
	if err != nil {
		log.Fatalln(err)
	}

	tbody, err := ioutil.ReadAll(trequest.Body)
	if err != nil {
		log.Fatalln(err)
	}

	telnet := string(tbody)
	fmt.Println(telnet)
}