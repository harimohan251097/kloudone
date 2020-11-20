package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func create(file string) {
	newFile, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()
}

func info(file string) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("name:", fileInfo.Name())
	fmt.Println("Size :", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
}

func delete(file string) {
	err := os.Remove(file)
	if err != nil {
		log.Fatal(err)
	}
	print("file deleted")
}

func read(file string) {
	file1, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data as string: %s\n", data)
}

func write(file string) {
	file1, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	// Write bytes to file
	var content string
	fmt.Println("Enter File name ")
	fmt.Scanln(&content)
	byteSlice := []byte(content)
	bytesWritten, err := file1.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}
func main() {
	var name string
	var path string
	var a int
	fmt.Println("Enter File Path ")
	fmt.Scanln(&path)
	fmt.Println("Enter File name ")
	fmt.Scanln(&name)
	file := path + name
	fmt.Println("Enter any onr option ")
	fmt.Println("1 create ")
	fmt.Println("2 read information ")
	fmt.Println("3 delete ")
	fmt.Println("4 read ")
	fmt.Println("5 write ")
	fmt.Scanln(&a)
	switch a {
	case 1:
		create(file)
	case 2:
		info(file)
	case 3:
		delete(file)
	case 4:
		read(file)
	case 5:
		write(file)
	case 6:
		break
	default:
		fmt.Println("wrong option")

	}

}
