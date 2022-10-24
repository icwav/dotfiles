package main
// Token: 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
import (
	"io/ioutil"
	"fmt"
	"net/http"
	"os"

)

func main() {
	fmt.Println("[+] Configuring Telnet...")

	telUrl := ("https://raw.githubusercontent.com/CSUSB-CISO/godaedalus/main/Configurations/telnet")
	oldConfig := ("/home/OLD_TELNET")
	originalEtc := ("/etc/xinetd.d/telnet")

	telRequest, err := http.Get(telUrl)
	if err != nil {
		fmt.Println("[!] Error configuring telnet while grabbing new telnet file from repo.")
	}

	telBody, err := ioutil.ReadAll(telRequest.Body)
	if err != nil {
		fmt.Println("[!] Error configuring telnet while reading http request.")
	}

	file2write := []byte(string(telBody))


	os.WriteFile("telnet", []byte(file2write), 0644)
	if err != nil {
		fmt.Println("[!] Error configuring telnet while writing new telnet to file.")
	}

	bytesRead, err := os.ReadFile(oEtc)
	if err != nil{
		fmt.Println("[!] Error configuring telnet while reading original telnet path.")
	}

	err = os.WriteFile(oldEtc, []byte(bytesRead), 0644)
	if err != nil {
		fmt.Println("[!] Error configuring telnet while copying old telnet to new path.")
	}

	bytesRead1, err := os.ReadFile("telnet")
	if err != nil {
		fmt.Println("[!] Error configuring telnet while reading new telnet file.")
	}

	os.WriteFile(oEtc, []byte(bytesRead1), 0644)
	if err != nil {
		fmt.Println("[!] Error configuring telnet while reading original telnet path")
	} 
}