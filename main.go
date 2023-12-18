package main

import (
	"DS2/Middleware"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	multicastAddress = "224.0.0.1"
	broadcastAddress = "0.0.0.0" // Empfang von allen IP-Adressen
	multicastPort    = 61424
)

func main() {
	var wg sync.WaitGroup

	person := Person{
		isScrumMaster: false,
		id:            uuid.New(),
	}
	// test: try to create a unicast listener
	unicastListener, err := Middleware.NewUnicastListener("localhost", 64399)
	unicastListener.Start()
	// try joining into room with given port
	broadcastSender, err := Middleware.NewBroadcastSender(broadcastAddress, multicastPort)
	broadcastSender.Send(fmt.Sprintf("joining: %s", person.id))
	// Signal-Handling, um die Anwendung sauber zu beenden
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Erstelle und starte den Multicast-Listener
	multicastListener, err := Middleware.NewMulticastListener(multicastAddress, multicastPort)
	if err != nil {
		log.Fatal("Fehler beim Starten des Multicast-Listeners:", err)
	}
	defer multicastListener.Close()

	// Erstelle und starte den Broadcast-Listener
	broadcastListener, err := Middleware.NewBroadcastListener(broadcastAddress, multicastPort)
	if err != nil {
		log.Fatal("Fehler beim Starten des Multicast-Listeners:", err)
	}
	defer broadcastListener.Close()

	wg.Add(1)
	go func() {
		defer wg.Done()
		broadcastListener.Listen()
	}()
	/*
		// Erstelle und starte den Multicast-Sender
		multicastSender, err := Middleware.NewMulticastSender(multicastAddress, multicastPort)
		if err != nil {
			log.Fatal("Fehler beim Starten des Multicast-Senders:", err)
		}
		defer multicastSender.Close()

		wg.Add(1)
		go func() {
			defer wg.Done()
			multicastSender.Send("Hallo, Multicast-Welt!")
		}()

		// Warte auf Signale zum Beenden der Anwendung
		<-signals

		// Schließe alle Ressourcen, wenn die Anwendung beendet wird
		multicastListener.Close()
		multicastSender.Close()
	*/
	// Warte auf das Beenden aller Goroutinen
	wg.Wait()

	fmt.Println("Anwendung wurde beendet.")
}
