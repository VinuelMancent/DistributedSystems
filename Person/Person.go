package Person

import "github.com/google/uuid"

type Person struct {
	Id            uuid.UUID
	IsScrumMaster bool
	TcpPorts      []int
}

func (p *Person) Elect() {

}
