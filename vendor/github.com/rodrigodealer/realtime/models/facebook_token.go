package models

import "encoding/json"

type FacebookToken struct {
	Token string `json:"access_token"`
	Type  string `json:"token_type"`
}

func (m *FacebookToken) FromJson(value string) error {
	err := json.Unmarshal([]byte(value), &m)
	return err
}
