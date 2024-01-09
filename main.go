package main

import (
	"bufio"
	"fmt"
	"note/note"
	"note/task"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	text := getTaskData()

	task, err := task.Create(text)

	if err != nil {
		fmt.Print(err)
		return
	}

	task.Display()

	err = saveData(task)

	if err != nil {
		return
	}

	note, err := note.Create(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}

	note.Display()

	err = saveData(note)

	if err != nil {
		return
	}
}

func getNoteData() (string, string) {
	noteTitle := getUserInput("Note Title:")
	noteContent := getUserInput("Note Content:")
	return noteTitle, noteContent
}

func getTaskData() string {
	return getUserInput("Task text:")
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving is failed")
		return err
	}

	fmt.Println("Saving is succeeded")

	return nil
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}