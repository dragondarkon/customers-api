package model

type Customer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
}
