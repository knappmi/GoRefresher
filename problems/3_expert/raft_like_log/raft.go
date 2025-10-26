package raftlog

// TODO: implement entries & replication

type Entry struct { Term int; Data []byte }

type Node struct { id int; log []Entry }

func (n *Node) Append(e Entry) { n.log = append(n.log, e) }
func (n *Node) LastTerm() int { if len(n.log)==0 { return 0 }; return n.log[len(n.log)-1].Term }
