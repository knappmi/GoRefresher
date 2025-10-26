package tableproperty

import (
	"testing"
)

func TestReverseTable(t *testing.T) {
	cases := []struct{ in, want string }{{"ab","ba"},{"racecar","racecar"},{"你好","好你"}}
	for _, c := range cases { if Reverse(c.in)!=c.want { t.Fatalf("%s => %s want %s", c.in, Reverse(c.in), c.want) } }
}

func TestReverseProperty(t *testing.T) {
	inputs := []string{"a","","abc","你好","🙂🙃"}
	for _, s := range inputs { if s != Reverse(Reverse(s)) { t.Fatalf("property fail %s", s) } }
}
