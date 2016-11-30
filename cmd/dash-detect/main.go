package main

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	if handle, err := pcap.OpenLive("en1", 1600, true, pcap.BlockForever); err != nil {
		panic(err)
	} else if err := handle.SetBPFFilter("arp"); err != nil {
		// optional
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			handlePacket(packet) // Do something with a packet here.
		}
	}
}

func getDeviceNames() ([]string) {
	result := make([]string, 5)

	interfaces, _ := net.Interfaces()
	for _, iface := range interfaces {
		append(result, iface.Name)
	}
	
	return result
}

func handlePacket(packet gopacket.Packet) {
	// if packet.LinkLayer().
	fmt.Println("Packet received", packet.Dump())
}
