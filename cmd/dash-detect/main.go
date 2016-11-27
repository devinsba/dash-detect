package main

import(
  "fmt"

  "github.com/google/gopacket"
  "github.com/google/gopacket/pcap"
)

func main() {
  if handle, err := pcap.OpenLive("eno1", 1600, true, pcap.BlockForever); err != nil {
    panic(err)
  } else if err := handle.SetBPFFilter("arp"); err != nil {  // optional
    panic(err)
  } else {
    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    for packet := range packetSource.Packets() {
      handlePacket(packet)  // Do something with a packet here.
    }
  }
}

func handlePacket(packet gopacket.Packet) {
  fmt.Println("Packet received", packet.Dump())
}
