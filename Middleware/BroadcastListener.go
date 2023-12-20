package Middleware

import (
	"fmt"
	"net"
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
func (bL *BroadcastListener) Listen() error {
	fmt.Println("Broadcast-Listener gestartet. Warte auf Nachrichten...")

	for {
		buffer := make([]byte, 1024)
		n, _, err := bL.conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Fehler beim Lesen der Nachricht:", err)
			return err
		}

		message := string(buffer[:n])
		if message == "join" {
			//send message back including all the room information
		} else if message == "Hello, UDP Broadcast!" {
			fmt.Println("Joa, das war n Hallo")
		}
		fmt.Printf("Empfangene Broadcast-Nachricht: %s\n", message)
	}
}

func (bL *BroadcastListener) ListenWithChannel(broadcastChannel chan string) error {
	fmt.Println("Broadcast-Listener mit channel gestartet. Warte auf Nachrichten...")

	for {
		buffer := make([]byte, 1024)
		n, _, err := bL.conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Fehler beim Lesen der Nachricht:", err)
			return err
		}
		message := string(buffer[:n])
		broadcastChannel <- message
	}
}

func (bL *BroadcastListener) ListenToChannelWithExpectedMessage(broadcastChannel chan string, expectedMessage string) error {
	for {
		buffer := make([]byte, 1024)
		n, _, err := bL.conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Fehler beim Lesen der Nachricht:", err)
			return err
		}
		message := string(buffer[:n])
		if message == expectedMessage {
			broadcastChannel <- message
		}
	}
}

// Close schließt den Broadcast-Listener.
func (bL *BroadcastListener) Close() {
	_ = bL.conn.Close()
}
