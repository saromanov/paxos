package paxos

import
(
	"time"
)


type Acceptor {
	starttime *time.Time
	messages []*Message
	number uint

}

func NewAcceptor()*Acceptor {
	acceptor := new(Acceptor)
	acceptor = time.Now()
	acceptor.messages = []*Message{}
	return acceptor
}

func (acceptor *Acceptor) Start() {
	go acceptor.process()
}

func (acceptor*Acceptor) process(){
	for {
		acceptor, err := acceptor.nextMessage()
		if err != nil {
			continue
		}

		switch acceptor.Type() {
			case prepare:
				if acceptor.Ballot() > acceptor.number {
					acceptor.number = acceptor.Ballot()
				}
			case promise:
				if acceptor.Ballot() == acceptor.number {

				}

			default:
				//TODO Send empty request
				continue
		}
	}
}

func (acceptor *Acceptor) nextMessage()(*Message, error) {
	if len(messages) == 0 {
		return nil, fmt.Errorf("%s", "queue is empty")
	}

	return messages[0], nil
}