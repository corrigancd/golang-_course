package ioexample

import (
	"fmt"
	"io"
	"os"
)

func ReadFileExample() {
	srcFile, err := os.Open("io/example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer srcFile.Close()

	fmt.Println("Reading file...")
	buf := make([]byte, 1024)
	for {
		n, err := srcFile.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading file:", err)
			break
		}
		if n == 0 {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}

func WriteFileExample() {
	destFile, err := os.Create("write_output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer destFile.Close()

	message := []byte("TEST CONTENT")

	fmt.Println("Writing to file...")
	_, err = destFile.Write(message)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func CopyFileExample() {
	srcFile, err := os.Open("io/example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer srcFile.Close()

	destFile, err := os.Create("copy_output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer destFile.Close()

	fmt.Println("Copying file...")
	written, err := io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println(fmt.Errorf("error copying file").Error())
	}

	fmt.Printf("written %v", written)
}
