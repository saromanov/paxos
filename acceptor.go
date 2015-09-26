package paxos

import
(
	"time"
)

type Acceptor {
	starttime *time.Time
}

func NewAcceptor()*Acceptor {
	acceptor := new(Acceptor)
	acceptor = time.Now()
}

func (acceptor*Acceptor) Process(){
	for {
		
	}
}