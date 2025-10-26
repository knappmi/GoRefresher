package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCRUD(t *testing.T) {
	st := newStore(); mux := http.NewServeMux(); RegisterHandlers(mux, st)
	ts := httptest.NewServer(mux); defer ts.Close()
	resp, err := http.Post(ts.URL+"/items", "application/json", strings.NewReader(`{"data":"hello"}`))
	if err!=nil { t.Fatal(err) }
	var it Item; json.NewDecoder(resp.Body).Decode(&it); if it.Data!="hello" { t.Fatal("create failed") }
	resp2, err := http.Get(ts.URL+"/items")
	if err!=nil { t.Fatal(err) }
	var list []Item; json.NewDecoder(resp2.Body).Decode(&list); if len(list)!=1 { t.Fatal("list failed") }
}
