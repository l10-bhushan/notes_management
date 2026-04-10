package model

type Error struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
