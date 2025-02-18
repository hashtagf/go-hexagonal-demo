package domains

import (
	"encoding/json"
	"errors"
)

type Order struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("id is required")
	}
	return nil
}

func (o *Order) ToJSON() string {
	json, err := json.Marshal(o)
	if err != nil {
		return ""
	}
	return string(json)
}
