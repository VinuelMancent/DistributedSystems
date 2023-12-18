package Middleware

import (
	"fmt"
	"net"
)

// BroadcastSender repräsentiert einen Broadcast-Sender.
type BroadcastSender struct {
	conn *net.UDPConn
}

// NewBroadcastSender erstellt einen neuen Broadcast-Sender.
func NewBroadcastSender(multicastAddress string, multicastPort int) (*BroadcastSender, error) {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", multicastAddress, multicastPort))
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}

	return &BroadcastSender{conn: conn}, nil
}

// Send sendet eine Broadcast-Nachricht.
func (bs *BroadcastSender) Send(message string) {
	_, err := bs.conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Fehler beim Senden der Broadcast-Nachricht:", err)
	}
}

// Close schließt den Broadcast-Sender.
func (bs *BroadcastSender) Close() {
	bs.conn.Close()
}
