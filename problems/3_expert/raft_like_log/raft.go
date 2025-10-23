package raftlog

// TODO: implement entries & replication

type Entry struct { Term int; Data []byte }

type Node struct { id int; log []Entry }

func (n *Node) Append(e Entry) { /* TODO */ }
