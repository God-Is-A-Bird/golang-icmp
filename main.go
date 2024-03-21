package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {

	type rawPacket struct {
		HEADER_IP    HEADER_IP    // 20 Bytes
		HEADER_ICMP  HEADER_ICMP  // 8 Bytes
		ICMP_PAYLOAD ICMP_PAYLOAD // Remaining Bytes
	}

	type HEADER_IP struct {
		VERSION       int // First 4 bits
		HEADER_LENGTH int // Second 4 bits
		SERVICE_TYPE
		TOTAL_LENGTH
		IDENTIFICATION
		FLAGS
		FRAGMENTATION_OFFSET
		TTL
		PROTOCOL
		CHECKSUM
	}

	type HEADER_ICMP struct {
		TYPE
		CODE
		CHECKSUM
		FIELD
	}

	type ICMP_PAYLOAD struct {
	}

	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		log.Print(err)
		return
	}

	for {

		maxAllowedSize := 1500
		buffer := make([]byte, 0, maxAllowedSize)
		bytesRecieved, err := syscall.Read(fd, buffer[len(buffer):cap(buffer)])
		buffer = buffer[:len(buffer)+bytesRecieved]
		if err != nil {
			log.Print(err)
			return
		}

		// Handle response when packet was too large
		if bytesRecieved >= maxAllowedSize {
			log.Print("Packet Trimmed!!!")
		}

		fmt.Print("\n Bytes Recieved: ", bytesRecieved, "\n", buffer)
		//fmt.Printf("\n%02X ", buffer)
	}

}
