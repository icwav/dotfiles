package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func prelimcheck() bool {
	if os.Getegid() != 0 {
		fmt.Println("You must run this as root")
		return false
		os.Exit(1)
	}
	return true
}

func main() {
	if prelimcheck() {
		fmt.Println("You are root, lets do some hardening!")
	}
	/*
		Look into namespace.conf to see if polyinstantiation is a good security move
		https://www.cyberciti.biz/faq/linux-namespaces-linux-kernel-virtualization/

		Install and configure fail2ban
		https://www.cyberciti.biz/faq/how-to-install-fail2ban-on-linux/

		GAIN NETWORK CONTROL set firewall rules
		https://www.cyberciti.biz/faq/linux-firewall-firewall-rules-linux-firewall-commands/

		check permissions for /etc/passwd && /etc/shadow
		https://www.cyberciti.biz/faq/linux-check-file-permissions-linux-command/

		find GTFObin suid and guid binaries and if so remove the suid and guid bit
		https://www.cyberciti.biz/faq/linux-find-suid-and-guid-files-linux-command/



		Disable wack http methods for apache
		https://www.cyberciti.biz/faq/linux-apache-disable-http-methods-apache-command/

		For SSH stuff:
		check permissions for /etc/ssh/ stuff
		https://dev-sec.io/baselines/ssh/

		import baseline secure sshd config file and overwrite sshd_config file in place


		Disable rsh
		https://www.cyberciti.biz/faq/linux-disable-rsh-rsh-server-linux-command/

		Disable rlogin
		https://www.cyberciti.biz/faq/linux-disable-rlogin-rlogin-server-linux-command/

		Disable rcp
		https://www.cyberciti.biz/faq/linux-disable-rcp-rcp-server-linux-command/

		https://www.cyberciti.biz/faq/linux-disable-rshd-rshd-server-linux-command/

		Disable rlogind
		https://www.cyberciti.biz/faq/linux-disable-rlogind-rlogind-server-linux-command/

		Disable rlogin
		https://www.cyberciti.biz/faq/linux-disable-rlogin-rlogin-server-linux-command/

		Limit failed login attempts
		https://www.cyberciti.biz/faq/linux-limit-failed-login-attempts-linux-command/

		Generate audit records for certain commands
		https://www.cyberciti.biz/faq/linux-audit-records-for-certain-commands-linux-command/




	*/

}

//possibly just remove these shit services: https://www.cyberciti.biz/tips/linux-security.html

func Telnet() {
	/*	err := os.Create("telnet")
		if err != nil {
			log.Fatal(err)
		}
	*/
	url := ("https://raw.githubusercontent.com/CSUSB-CISO/godaedalus/main/Configurations/telnet")
	
		trequest, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
	
		tbody, err := ioutil.ReadAll(trequest.Body)
		if err != nil {
			log.Fatalln(err)
		}
	
		file2write := []byte(string(tbody))
	
		/*what, err := os.Write(file2write)
		if err != nil {
			log.Fatalln(err)
		}
	*/
	
		os.WriteFile("telnet", []byte(file2write), 0644)
		if err != nil {
			log.Fatalln(err)
		}
	
	}


func disableCoreDumps() {
	fmt.Println("[+] Disabling core dumps...")
	_, err := exec.Command("echo", "'* hard core 0'", ">>", "/etc/security/limits.conf").Output()
	if err != nil {
		fmt.Println("[!] Error encountered while disabling core dumps: ", err)
	}
	_, err = exec.Command("echo", "'* soft core 0'", ">>", "/etc/security/limits.conf").Output()
	if err != nil {
		fmt.Println("[!] Error encountered while disabling core dumps: ", err)
	}
	_, err = exec.Command("echo", "'fs.suid_dumpable=0'", ">>", "/etc/sysctl.conf").Output()
	if err != nil {
		fmt.Println("[!] Error encountered while disabling core dumps: ", err)
	}
	// /etc/sysctl.d/9999-disable-core-dump.conf
	_, err = exec.Command("sudo", "sysctl", "-p", "/etc/sysctl.d/99-disable-core-dump.conf").Output()
	_, _ = exec.Command("sudo", "sysctl", "-p", "/etc/sysctl.d/9999-disable-core-dump.conf").Output()
	if err != nil {
		fmt.Println("[!] Error encountered while disabling core dumps: ", err)
	}
}

func accessRootLogin() {
	// access.conf
	fmt.Println("[+] Enabling root access from specific IP address...")
	_, err := exec.Command("echo", "+:root:192.168.89.1", ">>", "/etc/security/access.conf").Output()
	if err != nil {
		fmt.Println("[!] Error encountered while modifying access.conf: ", err)
	}
}

// Have to work on this to make it more robust, identify more package managers & ensure it is accurate
func identifyPackageManager() {
	fmt.Println("[+] Identifying package manager...")
	supportedPackMan := []string{"apt", "yum", "pacman"}
	for i:=0; i<len(supportedPackMan);i++ {
		out, err := exec.Command("which", supportedPackMan[i]).Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(out))
			//_, err = exec.Command("sudo", supportedPackMan[i], "install", "fail2ban").Output()
			//_, err = exec.Command("sudo", supportedPackMan[i], "install", "auditd", "audispd-plugins").Output()
		}
	}
}

func setLoginAttempts() {
	fmt.Println("[+] Setting up account lockout. After 3 login unsuccessful login attempts account will be locked for 10 minutes (root included)")
	lockoutCmd := "sudo echo auth    required           pam_tally2.so onerr=fail deny=3 unlock_time=600 audit even_deny_root root_unlock_time=600 >> /etc/pam.d/common-auth"
	_, err := exec.Command("bash", "-c", lockoutCmd).Output()
	if err != nil {
		fmt.Println("[-] Error encountered while setting up account lockout: ", err)
	}
}

func setTcpSynCookies() {
	fmt.Println("[+] Setting TCP Syn Cookies...")
	_, err := exec.Command("sudo", "echo", "'net.ipv4.tcp_syncookies = 1'", ">>", "/etc/sysctl.conf").Output()
	if err != nil {
		fmt.Println("[-] Error encountered while setting up TCP SYN Cookies:", err)
	}
	_, err = exec.Command("sudo", "sysctl", "-p").Output()
	if err != nil {
		fmt.Println(err)
	}
}

func disableIPv6() {
	fmt.Println("[+] Disabling IPv6...")
	_, err := exec.Command("sudo", "echo", "'net.ipv6.conf.all.disable_ipv6 = 1'", ">>", "/etc/sysctl.conf").Output()
	if err != nil {
		fmt.Println("[!] Error encountered while disabling IPv6: ", err)
	}
	_, err = exec.Command("sudo", "sysctl", "-p").Output()
	if err != nil {
		fmt.Println(err)
	}
}

func disableBlankPassAccts() {
	fmt.Println("[+] Disabling accounts with blank passwords...")
	findCmd := "sudo getent shadow | grep '^[^:]*::' | cut -d: -f1"
	out, err := exec.Command("bash", "-c", findCmd).Output()
	if err != nil {
		fmt.Println("[-] Error while finding accounts with blank passwords: ", err)
	}
	// iterate through each account w/blank password & disable it
	for _, line := range strings.Split(strings.TrimSuffix(string(out), "\n"), "\n") {
		//line -> username of account
		fmt.Println("[+] Disabling " + line + " since it has blank password...")
		disableCmd := "sudo usermod -L " + line
		_, err = exec.Command("bash", "-c", disableCmd).Output()
		if err != nil {
			fmt.Println("[-] Error encountered when disabling account ", line, err)
		}
	}
}

func findSUIDSGIDBits() {
	fmt.Println("[+] Finding SUID/SGID bits...")
	cmd := "sudo find / -type f \\( -perm -04000 -o -perm -02000 \\)"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println("[-] Error encountered while checking file permissions. Error: ", err)
	}
	fmt.Println("SUID/SGID bits: ", string(out))
}

// Can potentially send these permissions to web server & store it there
func checkFilePermissions() {
	fmt.Println("[+] Checking file permissions...")
	files := []string{"/var/log", "/etc/shadow", "/etc/passwd"}
	varLogCmd := "stat --printf='Permissions for %n are %A' "
	//iterate through each file, print permissions for files
	for _, file := range files {
		perms, err := exec.Command("bash", "-c", varLogCmd + file).Output()
		if err != nil {
			fmt.Println("[-] Error encountered while checking file permissions. Error: ", err)
		}
		fmt.Println(string(perms))
	}
}

//secureLocation -> full path of the secure sshd_config file
func sshdConfigReplacement(secureLocation string) {
	fmt.Println("[+] Replacing the sshd config file...")
	cmd := "cp " + secureLocation + "/etc/ssh/sshd_config"
	_, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println("[!] Error encountered while replacing sshd config file! Error: ", err)
	}
}

//disable rsh service
//filepath can either be /etc/xinet.d/rlogin or /etc/xinetd.d/rsh
//typical config file: https://ofstack.com/Linux/21673/install-rsh-under-ubuntu-for-passwordless-access.html
func disableRSH(filepath string) {
	ogConfigContent, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error encountered while disabling RSH: ", err)
	}
	newConfigFile := strings.Replace(string(ogConfigContent), "    disable = no", "    disable = yes", -1)
	// Golang will auto overwrite file if it already exists
	fileHandle, err := os.Create(filepath)
	if err != nil {
		fmt.Println("[!] Error encountered while replacing config file for RSH: ", err)
	}
	fileHandle.WriteString(string(newConfigFile))
	fileHandle.Close()
}
