package main

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	devices := getDeviceNames()

	for _, device := range devices {
		wg.Add(1)
		go captureDevice(wg, device)
	}
	wg.Wait()
}

func captureDevice(wg sync.WaitGroup, device string) {
	defer wg.Done()
	if handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever); err != nil {
		fmt.Print(device + " ")
		panic(err)
	} else if err := handle.SetBPFFilter("arp"); err != nil {
		fmt.Println(device + " error on setting filter")
		return
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			handlePacket(device, packet) // Do something with a packet here.
		}
	}
}

func getDeviceNames() ([]string) {
	var result []string

	interfaces, _ := net.Interfaces()
	for _, iface := range interfaces {
		result = append(result, iface.Name)
	}

	return result[:len(interfaces)]
}

func handlePacket(device string, packet gopacket.Packet) {
	// if packet.LinkLayer().
	fmt.Println("Packet received", device, packet.Dump())
}
