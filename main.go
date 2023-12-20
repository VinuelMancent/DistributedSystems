package main

import (
	"DS2/Middleware"
	"DS2/Person"
	"DS2/RoomState"
	"DS2/Ticket"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	multicastAddress          = "224.0.0.1"
	broadcastAddress          = "0.0.0.0" // Empfang von allen IP-Adressen
	multicastPort             = 61424
	timeToWaitForJoinResponse = 5 * time.Second
)

func main() {
	var wg sync.WaitGroup
	var roomState RoomState.RoomState

	joinRequestChannel := make(chan string)

	person := Person.Person{
		IsScrumMaster: false,
		Id:            uuid.New(),
		TcpPorts:      []int{},
	}

	// Erstelle und starte den Broadcast-Listener zum join request
	wg.Add(1)
	go func() {
		defer wg.Done()
		broadcastListener, err := Middleware.NewBroadcastListener(broadcastAddress, multicastPort)
		if err != nil {
			log.Fatal("Fehler beim Starten des Broadcast-Listeners:", err)
		}
		defer broadcastListener.Close()
		//err = broadcastListener.Listen()
		err = broadcastListener.ListenWithChannel(joinRequestChannel)
		if err != nil {
			log.Fatal("Fehler beim Zuhören auf BroadcastListener")
		}
	}()

	// try joining into room with given port
	broadcastSender, err := Middleware.NewBroadcastSender(broadcastAddress, multicastPort)
	if err != nil {
		fmt.Println(err.Error())
	}
	broadcastSender.Send(fmt.Sprintf("joining: %s", person.Id))
	// Höre auf eine Antwort.
	select {
	// Wenn eine Antwort im Zeitrahmen kommt, joine dem Raum
	case msg := <-joinRequestChannel:
		fmt.Printf("Got %s from worker\n", msg)
		_ = json.Unmarshal([]byte(msg), &roomState)
		fmt.Printf("Joined Room with %d members", len(roomState.Persons))
		break
	// Wenn keine Antwort kommt, erstelle den Raum und werde Verantwortlicher
	case <-time.After(timeToWaitForJoinResponse):
		fmt.Println("No Room availabe, creating a new Room")
		person.IsScrumMaster = true
		roomState = RoomState.RoomState{
			Persons:           []Person.Person{person},
			Tickets:           []Ticket.Ticket{},
			AvailableTcpPorts: []int{},
			CurrentPhase:      "Phase1",
		}
		break
	}
	//Create an UnicastListener in an own GoRoutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		unicastListener, err := Middleware.NewUnicastListener("localhost", 64399)
		if err != nil {
			fmt.Println(err.Error())
		}
		unicastListener.Start()
	}()

	// Signal-Handling, um die Anwendung sauber zu beenden
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer wg.Done()
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
