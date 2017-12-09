package models

import (
	"encoding/json"
)

type FacebookUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (m *FacebookUser) ToJson() string {
	userJSON, _ := json.Marshal(m)
	return string(userJSON)
}

func (m *FacebookUser) FromJson(value string) error {
	err := json.Unmarshal([]byte(value), &m)
	return err
}
