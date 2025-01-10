package main

import (
	"fmt"
	"net"
	"time"

	"github.com/Xsidelight/dns-resolver/internal/model"
)

func main() {
	header := model.Header{
		ID:    22,
		Flags: 0,
	}
	header.SetQR(false)
	header.SetRD(true)

	question := model.Question{
		Name:  "dns.google.com",
		Type:  1,
		Class: 1,
	}

	message := model.Message{
		Header:    header,
		Questions: []model.Question{question},
	}

	data, err := message.ToBytes()
	if err != nil {
		fmt.Println("Error creating DNS query:", err)
		return
	}

	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println("Error connecting to DNS server:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Error sending DNS query:", err)
		return
	}

	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	response := make([]byte, 512)
	n, err := conn.Read(response)
	if err != nil {
		fmt.Println("Error reading DNS response:", err)
		return
	}

	if response[0] != data[0] || response[1] != data[1] {
		fmt.Println("Response ID does not match query ID")
		return
	}

	fmt.Println("Received DNS response:", response[:n])
}
