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

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func main() {
	title, content := getNoteData()
	text := getTaskData()

	task, err := task.Create(text)

	if err != nil {
		fmt.Print(err)
		return
	}

	err = outputData(task)

	if err != nil {
		return
	}

	note, err := note.Create(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}

	err = outputData(note)

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

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
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