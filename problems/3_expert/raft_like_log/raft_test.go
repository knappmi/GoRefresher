package raftlog

import "testing"

func TestAppend(t *testing.T) {
	n := &Node{}
	n.Append(Entry{Term:1}); n.Append(Entry{Term:2})
	if n.LastTerm()!=2 || len(n.log)!=2 { t.Fatal("append failed") }
}
