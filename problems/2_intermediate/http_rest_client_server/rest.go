package rest

import (
	"errors"
	"net/http"
	"sync"
	"encoding/json"
	"time"
	"fmt"
)

type Item struct { ID int `json:"id"`; Data string `json:"data"` }

type store struct { mu sync.RWMutex; m map[int]Item; next int }
func newStore() *store { return &store{m: make(map[int]Item)} }

func (s *store) create(data string) Item { s.mu.Lock(); defer s.mu.Unlock(); s.next++; it := Item{ID:s.next, Data:data}; s.m[it.ID]=it; return it }
func (s *store) get(id int) (Item, bool) { s.mu.RLock(); defer s.mu.RUnlock(); it, ok := s.m[id]; return it, ok }
func (s *store) delete(id int) { s.mu.Lock(); defer s.mu.Unlock(); delete(s.m,id) }
func (s *store) list() []Item { s.mu.RLock(); defer s.mu.RUnlock(); out := make([]Item,0,len(s.m)); for _, v := range s.m { out = append(out,v) }; return out }

// RegisterHandlers attaches CRUD handlers.
func RegisterHandlers(mux *http.ServeMux, s *store) {
	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request){
		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(s.list())
		case http.MethodPost:
			var body struct{ Data string }
			if err := json.NewDecoder(r.Body).Decode(&body); err!=nil { http.Error(w, err.Error(), 400); return }
			it := s.create(body.Data)
			json.NewEncoder(w).Encode(it)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/items/", func(w http.ResponseWriter, r *http.Request){
		var id int
		_, err := fmt.Sscanf(r.URL.Path, "/items/%d", &id)
		if err!=nil { http.Error(w, "bad id", 400); return }
		switch r.Method {
		case http.MethodGet:
			it, ok := s.get(id); if !ok { w.WriteHeader(404); return }; json.NewEncoder(w).Encode(it)
		case http.MethodDelete:
			s.delete(id); w.WriteHeader(204)
		default:
			w.WriteHeader(405)
		}
	})
}

// RetryClient performs GET with simple exponential backoff.
func RetryClient(url string, attempts int) (*http.Response, error) {
	var err error
	for i:=0; i<attempts; i++ {
		resp, e := http.Get(url)
		if e==nil { return resp, nil }
		err = e
		time.Sleep(time.Duration(1<<i) * 10 * time.Millisecond)
	}
	return nil, err
}

var ErrMaxAttempts = errors.New("max attempts")
