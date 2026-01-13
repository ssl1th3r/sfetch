package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	emoticons := []string{"<3", ":3", ":D", ":)", ";)", ":P", ":*", ":o", ":/", ":|", ":'(", "^_^", ">_<", "o.o", "-_-", ">:3", "<:p", "XD", "O-O", ">:P"}
	rand.Seed(time.Now().UnixNano())

	OSS := exec.Command("sh", "-c", "tr -d '\"' < /etc/os-release | grep PRETTY_NAME | cut -b 13-")
	kernel := exec.Command("uname", "-sr")
	shell := os.Getenv("SHELL")
	upt := exec.Command("uptime")
	wm := os.Getenv("DESKTOP_SESSION")
	dm := exec.Command("/bin/sh", "-c", "systemctl status display-manager | grep -oP '(?<=/)\\w+(?=\\.service)' | head -1")

	outputos, _ := OSS.CombinedOutput()
	outputkr, _ := kernel.CombinedOutput()
	outputuptime, _ := upt.CombinedOutput()
	outputdm, _ := dm.CombinedOutput()

	fmt.Printf("------------ \n")
	fmt.Printf("OS: %s", outputos)
	fmt.Printf("KERNEL: %s", outputkr)
	fmt.Printf("SHELL: %s", shell+"\n")
	fmt.Printf("UPTIME: %s", outputuptime)
	fmt.Printf("WM: %s", wm+"\n")
	fmt.Printf("DM: %s", outputdm)
	fmt.Printf("%s\n", emoticons[rand.Intn(len(emoticons))])
	fmt.Printf("------------\n")

}
