package zeroallocjson

import (
	"reflect"
	"sync"
)

var bufPool = sync.Pool{ New: func() any { b := make([]byte, 0, 256); return &b } }

// Marshal marshals a struct with string/int/bool fields (simplified) avoiding extra allocations.
func Marshal(v any) []byte {
	rv := reflect.ValueOf(v)
	if rv.Kind()==reflect.Pointer { rv = rv.Elem() }
	if rv.Kind()!=reflect.Struct { return []byte("null") }
	bptr := bufPool.Get().(*[]byte)
	buf := (*bptr)[:0]
	buf = append(buf, '{')
	rt := rv.Type()
	for i:=0;i<rv.NumField();i++ {
		if i>0 { buf = append(buf, ',') }
		fname := rt.Field(i).Name
		buf = append(buf, '"')
		buf = append(buf, fname...)
		buf = append(buf, '"', ':')
		fv := rv.Field(i)
		switch fv.Kind() {
		case reflect.String:
			buf = append(buf, '"'); buf = append(buf, fv.String()...); buf = append(buf, '"')
		case reflect.Int, reflect.Int64, reflect.Int32:
			buf = append(buf, intToBytes(fv.Int())...)
		case reflect.Bool:
			if fv.Bool() { buf = append(buf, 't','r','u','e') } else { buf = append(buf, 'f','a','l','s','e') }
		default:
			buf = append(buf, '"','?','"')
		}
	}
	buf = append(buf, '}')
	out := make([]byte, len(buf))
	copy(out, buf)
	bufPool.Put(bptr)
	return out
}

func intToBytes(n int64) []byte { if n==0 { return []byte{'0'} }; neg := n<0; if neg { n=-n }; var tmp [20]byte; i:=len(tmp); for n>0 { i--; tmp[i] = byte('0'+n%10); n/=10 }; if neg { i--; tmp[i]='-' }; return tmp[i:] }
