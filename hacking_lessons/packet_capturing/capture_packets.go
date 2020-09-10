package main

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"log"
)

var (
	iface = "wlan0"
	devFound = false
	snaplen = int32(1600)
	timeout = pcap.BlockForever
	filter = "tcp and port 443"
	promisc = false
) 

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("Finding devices")
	for _, dev := range devices {
		if dev.Name == iface{
			devFound = true
		}
		if !devFound {
			log.Panicf("Device not found %s\n", iface)
		}
		handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
		if err != nil {
			log.Panicln(err)
		}

		defer handle.Close()

		err := handle.SetBPFFilter(filter)
		
		if 
	}
}
