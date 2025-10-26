package interfacereflection

import "testing"

func TestDescribeStruct(t *testing.T) {
	type X struct { A int; B string }
	names := DescribeStruct(X{})
	if len(names)!=2 || names[0]!="A" || names[1]!="B" { t.Fatalf("bad names %#v", names) }
}
