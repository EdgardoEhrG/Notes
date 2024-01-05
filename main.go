package main

import (
	"fmt"
	"note/note"
)

func main() {
	title, content := getNoteData()

	note, err := note.CreateNewNote(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}
}

func getNoteData() (string, string) {
	noteTitle := getUserInput("Note Title:")
	noteContent := getUserInput("Note Content:")

	return noteTitle, noteContent
}

func getUserInput(prompt string) string {
	var value string
	var valuePointer *string = &value

	fmt.Print(prompt)

	fmt.Scan(valuePointer)

	return value
}