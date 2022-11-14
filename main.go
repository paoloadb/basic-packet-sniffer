package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	devName := GetDevices()
	fmt.Println("Starting Capture (Ctrl+C to Stop)")
	handle, err := pcap.OpenLive(devName, 2048, true, pcap.BlockForever)
	checkError(err)
	defer handle.Close()

	// err = handle.SetBPFFilter("dst net 192.168.254.107")
	// checkError(err)

	source := gopacket.NewPacketSource(handle, handle.LinkType())

	for packets := range source.Packets() {
		pkt := packets.ApplicationLayer()
		if pkt != nil {
			fmt.Println(packets)
			
			// dstIP := packets.NetworkLayer().NetworkFlow().Dst()
			// srcIP := packets.NetworkLayer().NetworkFlow().Src()
			// x := packets.TransportLayer().TransportFlow()
			// fmt.Println("Source IP", srcIP)
			// fmt.Println("Destination IP", dstIP)
			// fmt.Println("Source Port",x.Src(), "Destination Port", x.Dst())
			// fmt.Println(string(pkt.Payload()))
		}
		
	}
}

func GetDevices() string {
	var devNum int 
	
	fmt.Println("Retrieving all capture devices...")
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalln("Unable to retrieve devices!")
	}

	for i, device := range devices {
		fmt.Println(i + 1, ">>", device.Name, ">>", device.Description)
		}
	fmt.Println("Select Capture Device Number:")
	fmt.Scanln(&devNum)
	return devices[devNum - 1].Name 
}