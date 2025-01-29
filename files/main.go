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
	readByScanner(path)
	remove(path)
	writeByteFile()
	/*
		> writeStringFile
		file /tmp/file.txt created! Size: 36

		> readAllContent
		Lorem ipsum
		Lorem ipsum
		Lorem ipsum

		> readByChunk
		(3) Lor
		(3) em
		(3) ips
		(3) um

		(3) Lor
		(3) em
		(3) ips
		(3) um

		(3) Lor
		(3) em
		(3) ips
		(3) um

		> readByScanner
		Lorem ipsum
		Lorem ipsum
		Lorem ipsum

		> remove
		file /tmp/file.txt removed

		> writeByteFile
		file created! Size: 11
	*/
}

func writeStringFile(path string) {
	fmt.Println("> writeStringFile")

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	size, err := f.WriteString(createContent(3))
	if err != nil {
		panic(err)
	}
	fmt.Printf("file %v created! Size: %d\n", path, size)
	// f.Close() // where defer is executed
}

func writeByteFile() {
	fmt.Println("> writeByteFile")

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
	fmt.Println("> readAllContent")

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))
}

func readByChunk(path string) {
	fmt.Println("> readByChunk")

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

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

func readByScanner(path string) {
	fmt.Println("> readByScanner")

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func remove(path string) {
	fmt.Println("> remove")

	err := os.Remove(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("file %v removed\n", path)
}

func createContent(numberOfLine uint) string {
	if numberOfLine == 0 {
		numberOfLine = 1
	}

	var value string
	content := "Lorem ipsum\n"

	for i := 0; i < int(numberOfLine); i++ {
		value = value + content
	}

	return value
}
