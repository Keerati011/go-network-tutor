package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the server ที่ localhost port 5000
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}

	//  ตัดการเชื่อมต่อเมื่อ ตัวแปร conn ปิดทำงาน
	defer conn.Close()

	fmt.Println("Connected to server")

	reader := bufio.NewReader(os.Stdin)
	for {
		// Read user input
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')

		// Send the message to the server
		conn.Write([]byte(message))

		// Print the number of bytes sent
		fmt.Printf("Sent %d bytes\n", len(message))

		// Receive and print the server's response
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		fmt.Printf("Server response: %s", buffer[:n])
	}
}
