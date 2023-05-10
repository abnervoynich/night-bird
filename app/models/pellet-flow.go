package models

type PelletFlow struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ProjectName string `json:"project-name"`
	Flow        []Flow `json:"flow"`
	IsActive    bool   `json:"is-active" `
}

type Flow struct {
	OrderID  int `json:"order-id"`
	PelletID int `json:"pellet-id"`
}
