package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	telnet, err := http.Get("https://raw.githubusercontent.com/CSUSB-CISO/godaedalus/main/Configurations/telnet?token=GHSAT0AAAAAABZQGHPBLQZAXJDZ3N2FCTSSY2DR3ZQ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(telnet)
}