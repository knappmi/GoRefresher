package zeroallocjson

import "testing"

type sample struct { Name string; Age int; Active bool }

func TestMarshal(t *testing.T) {
	b := Marshal(sample{Name:"A", Age:5, Active:true})
	if string(b) != "{\"Name\":\"A\",\"Age\":5,\"Active\":true}" { t.Fatalf("bad json %s", string(b)) }
}
