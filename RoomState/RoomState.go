package RoomState

import (
	"DS2/Person"
	"DS2/Ticket"
)

type RoomState struct {
	Persons           []Person.Person
	Tickets           []Ticket.Ticket
	AvailableTcpPorts []int
	CurrentPhase      string
}

// GetFreeTcpPort returns the first element in the slice and removes it
func (rs *RoomState) GetFreeTcpPort() int {
	elem := rs.AvailableTcpPorts[0]
	rs.AvailableTcpPorts = append(rs.AvailableTcpPorts[:0], rs.AvailableTcpPorts[0+1:]...)
	return elem
}

func (rs *RoomState) AddTcpPortsFromPersonToPool(person Person.Person) {
	rs.AvailableTcpPorts = append(rs.AvailableTcpPorts, person.TcpPorts...)
}
