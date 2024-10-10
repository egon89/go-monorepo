package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("> Files")
	path := "/tmp/file.txt"
	writeStringFile(path)
	readAllContent(path)
	readByChunk(path)
	remove(path)
	writeByteFile()
	/*
		file /tmp/file.txt created! Size: 11
		Lorem ipsum
		(3) Lor
		(3) em
		(3) ips
		(2) um
		file /tmp/file.txt removed
		file created! Size: 11
	*/
}

func writeStringFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	size, err := f.WriteString("Lorem ipsum")
	if err != nil {
		panic(err)
	}
	fmt.Printf("file %v created! Size: %d\n", path, size)
	// f.Close() // where defer is executed
}

func writeByteFile() {
	f, err := os.Create("/tmp/file-byte.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	size, err := f.Write([]byte("Lorem ipsum"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("file created! Size: %d\n", size)
	// f.Close() // where defer is executed
}

func readAllContent(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))
}

func readByChunk(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Printf("(%v) %v\n", n, string(buffer[:n]))
	}
}

func remove(path string) {
	err := os.Remove(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("file %v removed\n", path)
}
