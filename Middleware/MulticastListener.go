package Middleware

import (
	"fmt"
	"net"
)

// MulticastListener repräsentiert einen Multicast-Listener.
type MulticastListener struct {
	conn *net.UDPConn
}

// NewMulticastListener erstellt einen neuen Multicast-Listener.
func NewMulticastListener(multicastAddress string, multicastPort int) (*MulticastListener, error) {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", multicastAddress, multicastPort))
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}

	return &MulticastListener{conn: conn}, nil
}

// Listen lauscht auf eingehende Multicast-Nachrichten.
func (ml *MulticastListener) Listen() {
	fmt.Println("Multicast-Listener gestartet. Warte auf Nachrichten...")

	for {
		buffer := make([]byte, 1024)
		n, _, err := ml.conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Fehler beim Lesen der Nachricht:", err)
			return
		}

		message := string(buffer[:n])
		fmt.Printf("Empfangene Multicast-Nachricht: %s\n", message)
	}
}

// Close schließt den Multicast-Listener.
func (ml *MulticastListener) Close() {
	ml.conn.Close()
}
