package Middleware

import (
	"fmt"
	"net"
)

func sendMessage(message string, address string) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Beispiel: Sende Daten zum Server und empfange die Antwort
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}

	// Lese die Antwort vom Server
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Println("Received from server:", string(buffer[:n]))
}
