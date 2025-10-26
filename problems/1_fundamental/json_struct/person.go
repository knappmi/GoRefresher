package jsonstruct

import (
	"encoding/json"
	"time"
)

type RFC3339Time struct {
	time.Time
}

func (rt RFC3339Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(rt.Time.Format(time.RFC3339))
}

func (rt *RFC3339Time) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	tm, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}
	rt.Time = tm
	return nil
}

type Person struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Email     string      `json:"email,omitempty"`
	CreatedAt RFC3339Time `json:"created_at"`
}

// MarshalPerson marshals a Person.
func MarshalPerson(p Person) ([]byte, error) {
	return json.Marshal(p)
}

// UnmarshalPerson unmarshals bytes into Person.
func UnmarshalPerson(b []byte) (Person, error) {
	var p Person
	err := json.Unmarshal(b, &p)
	return p, err
}
