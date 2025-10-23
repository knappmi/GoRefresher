package jsonstruct

import (
	"encoding/json"
	"time"
)

type Person struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// MarshalPerson marshals a Person.
func MarshalPerson(p Person) ([]byte, error) {
	// TODO: customize time format / add validation
	return json.Marshal(p)
}

// UnmarshalPerson unmarshals bytes into Person.
func UnmarshalPerson(b []byte) (Person, error) {
	// TODO: custom parsing / validation
	var p Person
	err := json.Unmarshal(b, &p)
	return p, err
}
