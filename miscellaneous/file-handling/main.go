package main

import (
	"bufio"
	"fmt"
	"os"
)

const myFile string = "my-file.txt"

func main() {

	//remove file
	err := os.Remove(myFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		} else {
			panic(err)
		}
	}

	//create a new file
	file, err := os.Create(myFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//writing in new file
	// size, err := file.WriteString("capivara world!!!") //if file is only text you can use this other function
	size, err := file.Write([]byte("capivara world!!!"))
	if err != nil {
		panic(err)
	}

	//get content in file
	content, err := os.ReadFile(myFile)
	if err != nil {
		panic(err)
	}

	//To read big files
	// ReadInChuncks(myFile)

	fmt.Printf("Content in file: %s\n", string(content))
	fmt.Printf("Size file %d bytes \n", size)

}

func ReadInChuncks(fileName string) {
	content, err := os.Open(myFile)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(content)
	buf := make([]byte, 5)

	for {
		n, err := reader.Read(buf)
		if err != nil {
			break
		}

		fmt.Printf("Size: %d bytes\n", n)
		fmt.Printf("Content: %s\n", buf[:n])
	}
}
