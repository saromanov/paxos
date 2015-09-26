package paxos

import
(
	"net"
)

type Paxos struct {
	nodes []string
}

func (paxos*Paxos) AddNode(title string){
	paxos.nodes = append(paxos.nodes, title)
}