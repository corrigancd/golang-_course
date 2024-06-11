package main

import (
	"bufio"
	"example.com/note/note"
	"fmt"
	"os"
	"strings"
)

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func main() {
	title, content := getNoteData()

	userNote, noteErr := note.New(title, content)

	if nil != noteErr {
		fmt.Println(noteErr)
		return
	}

	userNote.Display()
	noteErr = userNote.Save()

	if noteErr != nil {
		fmt.Println("Saving the note failed")
		return
	}

	fmt.Println("Saving the note succeeded")
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
