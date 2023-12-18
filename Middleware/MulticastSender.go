package Middleware

import (
	"fmt"
	"net"
)

// MulticastSender repräsentiert einen Multicast-Sender.
type MulticastSender struct {
	conn *net.UDPConn
}

// NewMulticastSender erstellt einen neuen Multicast-Sender.
func NewMulticastSender(multicastAddress string, multicastPort int) (*MulticastSender, error) {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", multicastAddress, multicastPort))
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}

	return &MulticastSender{conn: conn}, nil
}

// Send sendet eine Multicast-Nachricht.
func (ms *MulticastSender) Send(message string) {
	_, err := ms.conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Fehler beim Senden der Multicast-Nachricht:", err)
	}
}

// Close schließt den Multicast-Sender.
func (ms *MulticastSender) Close() {
	ms.conn.Close()
}
