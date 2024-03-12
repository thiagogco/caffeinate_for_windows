package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("powershell", "powercfg", "/change", "standby-timeout-ac", "0;", "powershell", "powercfg", "/change", "monitor-timeout-ac", "0")

	_, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("windows is now Caffeinate")

}
