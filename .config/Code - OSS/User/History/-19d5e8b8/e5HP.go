package main
// Token: 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
import (
	"io/ioutil"
//	"log"
	"fmt"
	"net/http"
	"os"

)

func main() {

	telUrl := ("https://raw.githubusercontent.com/CSUSB-CISO/godaedalus/main/Configurations/telnet")
	//newetc := ("/home/ivan/CCDC/testdir/newconfigs/telnet")
	oldEtc := ("/home/ivan/CCDC/testdir/oldconfigs/telnet")
	oEtc := ("/home/ivan/CCDC/OLD_TELNET")

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

	fmt.Println("Configuring telnet...")
	os.WriteFile(oEtc, []byte(bytesRead1), 0644)
	if err != nil {
		fmt.Println("[!] Error configuring telnet while reading original telnet path")
	} 
	fmt.Println(err)
}