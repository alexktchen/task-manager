package models

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Status_bool bool   `json:"-"`
	Status      int    `json:"status"`
}
