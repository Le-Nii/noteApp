package main

import (
	"bufio"
	"fmt"
	"noteApp/note"
	"noteApp/todo"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo Text: ")

	userTodo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userTodo.Display()
	err = userTodo.Save()

	if err != nil {
		fmt.Println("Saving The todo Failed.")
	}
	fmt.Println("Saved the todo Successfully.")

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving The Note Failed.")
	}
	fmt.Println("Saved the Note Successfully.")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
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
