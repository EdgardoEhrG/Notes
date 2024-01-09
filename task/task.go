package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Task struct {
	Text string `json:"text"`
}

func (task Task) Display() {
	fmt.Println(task.Text)
}

func (task Task) Save() error {
	fileName := "task.json"

	json, err := json.Marshal(task)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func Create(text string) (Task, error) {

	if text == "" {
		return Task{}, errors.New("Invalid input")
	}

	return Task{
		Text: text,
	}, nil
}