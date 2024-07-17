package main

import (
	"bufio"
	"fmt"
	"noteApp/note"
	"noteApp/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

//type displayer interface {
//	Display()
//}

type outputtable interface {
	saver
	Display()
}

//type outputtable interface {
//	Save() error
//	Display()
//}

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

	err = outputData(userTodo)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}

}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving The Note Failed.")
		return err
	}
	fmt.Println("Saved the Note Successfully.")
	return nil

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
