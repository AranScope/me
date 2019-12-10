package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type Device struct {
	IpAddr string
	Name   string
}

func main() {
	cmd := exec.Command("nmap", "-sn", "-oG", "-", "192.168.1.0/24")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(output), "\n")
	lines = lines[1 : len(lines)-2]

	var devices []Device

	for _, line := range lines {
		words := strings.Split(strings.Split(line, "\t")[0], " ")
		name := strings.Split(words[2], ".")[0]
		name = strings.TrimPrefix(name, "(")
		name = strings.TrimSuffix(name, ")")

		device := Device{
			IpAddr: words[1],
			Name:   name,
		}

		devices = append(devices, device)
		fmt.Printf("%v\n", device)
	}
}
