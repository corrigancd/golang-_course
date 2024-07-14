package main

import (
	"bufio"
	"example.com/note/note"
	"example.com/note/todo"
	"fmt"
	"os"
	"strings"
)

// example of interface
type saver interface {
	Save() error
}

type outputtable interface {
	saver // example of embedded interface
	Display()
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving failed")
		return err
	}

	fmt.Println("Saving succeeded")
	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func main() {
	title, content := getNoteData()
	todoMessage := getUserInput("Todo Message:")

	userNote, err := note.New(title, content)
	if nil != err {
		return
	}

	todo, err := todo.New(todoMessage)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}

	outputData(todo)
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	// This won't work if you are using spaces
	// fmt.Scanln(&value)

	// Alternative
	reader := bufio.NewReader(os.Stdin)  // create a reader that listens to the command line
	text, err := reader.ReadString('\n') // read the command line until the first new line character

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n") // remove newline from text
	text = strings.TrimSuffix(text, "\r") // for windows newline also contains \r

	return text
}

// how to any value!! :) :)
func printSomething(value any) {
	typedValue, isInteger := value.(int) // testing the type of a variable
	if isInteger {
		fmt.Println(typedValue)
	}

	switch value.(type) { // how to retrieve the type of the value, this only works inside of a switch statement!
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Println("unknown")
	}
}

// example of generics
// [T int | float64 | string] means that T can be any type as long as it is int, float64 or string
func add[T int | float64 | string](a, b T) T {
	return a + b
}
