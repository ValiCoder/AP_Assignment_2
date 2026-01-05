package model

type Task struct {
	ID      string `json:"id"`
	Payload string `json:"payload"`
	Status  string `json:"status"`
}

func CreateTask() {

}
