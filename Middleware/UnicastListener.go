package Middleware

import (
	"fmt"
	"net"
	"os"
)

// UnicastListener repräsentiert einen Unicast-Server.
type UnicastListener struct {
	listener net.Listener
}

// NewUnicastListener erstellt einen neuen Unicast-Server.
func NewUnicastListener(address string, port int) (*UnicastListener, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		return nil, err
	}
	return &UnicastListener{listener: listener}, nil
}

// Start startet den Unicast-Server und wartet auf eingehende Verbindungen.
func (uS *UnicastListener) Start() {
	fmt.Println("Unicast-Server gestartet. Warte auf Verbindungen...")
	for {
		conn, err := uS.listener.Accept()
		if err != nil {
			fmt.Println("Fehler beim Akzeptieren der Verbindung:", err)
			os.Exit(1)
		}
		go uS.handleClient(conn)
	}
}

// handleClient behandelt die Kommunikation mit einem Client.
func (uS *UnicastListener) handleClient(conn net.Conn) {
	defer conn.Close()
	// Handle client communication here
	fmt.Println("Akzeptierte Verbindung von", conn.RemoteAddr())
	// Beispiel: Echo-Server
	buffer := make([]byte, 1024)
	for {
		// Lese Daten vom Client
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Fehler beim Lesen:", err)
			return
		}
		message := string(buffer[:n])
		fmt.Printf("Empfangene Unicast-Nachricht: %s\n", message)

		// Sende die Daten zurück zum Client
		_, err = conn.Write(buffer[:n])
		if err != nil {
			fmt.Println("Error writing:", err)
			return
		}
	}

}

// Stop stoppt den Unicast-Server und schließt die Verbindung.
func (uS *UnicastListener) Stop() {
	uS.listener.Close()
	fmt.Println("Unicast-Listener gestoppt.")
}
