package controller

import (
	"encoding/xml"
	"fmt"
	"github.com/AranScope/me/common/metrics"
	"github.com/AranScope/me/common/timeutils"
	"os/exec"
	"time"
)

type Device struct {
	IpAddr  string `json:"ip_addr"`
	MacAddr string `json:"mac_addr,omitempty"`
	Name    string `json:"name"`
	Vendor  string `json:"vendor"`
}

type NmapResponse struct {
	Hosts []struct {
		Addresses []struct {
			Addr   string `xml:"addr,attr"`
			Type   string `xml:"addrtype,attr"`
			Vendor string `xml:"vendor,attr"`
		} `xml:"address"`
		HostNames []struct {
			Name string `xml:"name,attr"`
		} `xml:"hostnames>hostname"`
	} `xml:"host"`
}

func (n *NmapResponse) toDevices() []Device {
	var devices []Device

	for _, host := range n.Hosts {
		device := Device{}
		for _, hostname := range host.HostNames {
			device.Name = hostname.Name
			break
		}

		for _, addr := range host.Addresses {
			if addr.Vendor != "" {
				device.Vendor = addr.Vendor
			}
			switch addr.Type {
			case "ipv4":
				{
					device.IpAddr = addr.Addr
					break
				}
			case "mac":
				device.MacAddr = addr.Addr
				break
			}
		}
		devices = append(devices, device)
	}

	return devices
}

var DeviceRegistry = map[string]Device{}

const scanInterval = time.Hour

func findDevices() ([]Device, error) {
	cmd := exec.Command("nmap", "-sn", "-oX", "-", "192.168.1.0/24")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	rsp := NmapResponse{}
	err = xml.Unmarshal(output, &rsp)
	if err != nil {
		return nil, err
	}

	return rsp.toDevices(), nil
}

func Init() {
	timeutils.Every(scanInterval, func() {
		fmt.Print("🤖 Updating device registry...\n")
		devices, err := findDevices()
		if err != nil {
			panic(err)
		}

		for _, device := range devices {
			DeviceRegistry[device.MacAddr] = device
			DeviceRegistry[device.Name] = device
		}
		fmt.Print("✅ Finished updating\n")

		metrics.Count("device_registry_updated", 1)
	})
}
