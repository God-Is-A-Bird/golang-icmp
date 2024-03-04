package main

import (
	"fmt"
	"net"
)

func echoListen(address, network string) {
	// Only accepts 64 byte ICMP, I don't understand why the total size ends up being 84 tho,
	// On an airplane right now so I'll check all this later
	// This listener will print both the echo request (code 8) and the echo response (code 0)
	payload := make([]byte, 84)
	request := 1
	response := 1
	listen, err := net.ListenPacket(network, address)
	if err != nil {
		fmt.Print("Error!!!")
	}
	for {
		listen.ReadFrom(payload)

		switch payload[0] {
		case 8:
			fmt.Print("System Is Making Echo Request ", request)
			request++
		case 0:
			fmt.Print("System Is Sending Echo Response ", response)
			response++
		default:
			fmt.Print("ICMP Packet of unsupported code recieved. Only codes 0 and 8 are currently supported")
		}

		fmt.Print("\nPayload: ", payload, "\n\n")
	}
}

func main() {
	echoListen("127.0.0.1", "ip4:icmp")
}
