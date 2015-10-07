package paxos

const (
	prepare = "prepare"
	promise = "promise"
)

type Message struct {
	messageType string
	number uint
}


func (m *Message) Type() {
	return m.nessageType
}

func (m *Message) Ballot() uint {
	return m.number
}