package models

type HealthcheckStatus struct {
	Status   string                `json:"status"`
	Services []HealthcheckServices `json:"services"`
}

type HealthcheckServices struct {
	Name  string `json:"name"`
	State string `json:"state"`
	Code  int    `json:"code"`
}
