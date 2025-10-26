package jsonstruct

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestRoundTrip(t *testing.T) {
	p := Person{ID:1, Name:"Alice", Email:"", CreatedAt: RFC3339Time{Time: time.Unix(0,0).UTC()}}
	b, err := MarshalPerson(p); if err!=nil { t.Fatalf("marshal: %v", err) }
	if strings.Contains(string(b), "email") { t.Fatalf("email field should be omitted: %s", string(b)) }
	p2, err := UnmarshalPerson(b); if err!=nil { t.Fatalf("unmarshal: %v", err) }
	if p2.ID!=p.ID || p2.Name!=p.Name || !p2.CreatedAt.Time.Equal(p.CreatedAt.Time) { t.Fatalf("mismatch: %#v vs %#v", p, p2) }
}

func TestTimeFormat(t *testing.T) {
	p := Person{CreatedAt: RFC3339Time{Time: time.Date(2024,1,2,3,4,5,0,time.UTC)}}
	b, _ := json.Marshal(p)
	if !strings.Contains(string(b), "2024-01-02T03:04:05Z") { t.Fatalf("bad time format: %s", string(b)) }
}
