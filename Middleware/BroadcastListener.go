package Middleware

import (
	"fmt"
	"net"
	"os"
)

// BroadcastListener repräsentiert einen Broadcast-Listener.
type BroadcastListener struct {
	conn *net.UDPConn
}

// NewBroadcastListener erstellt einen neuen Broadcast-Listener.
func NewBroadcastListener(broadcastAddress string, broadcastPort int) (*BroadcastListener, error) {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", broadcastAddress, broadcastPort))
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}

	return &BroadcastListener{conn: conn}, nil
}

// Listen lauscht auf eingehende Broadcast-Nachrichten.
func (bL *BroadcastListener) Listen() {
	fmt.Println("Multicast-Listener gestartet. Warte auf Nachrichten...")

	for {
		buffer := make([]byte, 1024)
		n, _, err := bL.conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Fehler beim Lesen der Nachricht:", err)
			os.Exit(1)
		}

		message := string(buffer[:n])
		fmt.Printf("Empfangene Broadcast-Nachricht: %s\n", message)
	}
}

// Close schließt den Broadcast-Listener.
func (bL *BroadcastListener) Close() {
	bL.conn.Close()
}
