package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type User struct {
	Id    string
	Email string
	Name  string
	Phone string
}

func main() {
	readAllFile()
	readLineByLine()
}

func readAllFile() {
	fmt.Println("> reading all file")

	path := "./resource/users_1k.csv"

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	if len(records) < 2 {
		panic(fmt.Errorf("no data found in the %v file", path))
	}

	headers := records[0]
	dataRows := records[1:]

	fmt.Println("Headers", headers)

	var users []User
	for _, row := range dataRows {
		if len(row) != 4 {
			fmt.Println("skipping malformed row:", row)
			continue
		}

		user := User{
			Id:    row[0],
			Email: row[1],
			Name:  row[2],
			Phone: row[3],
		}

		users = append(users, user)
	}

	for _, user := range users {
		fmt.Printf("User: %+v\n", user) // %+v prints field:value
	}
}

func readLineByLine() {
	fmt.Println("> reading line by line")

	path := "./resource/users_10k.csv"

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {
		panic(err)
	}
	fmt.Println("Headers:", header)

	for {
		row, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading row:", err)
			continue
		}

		if len(row) != 4 {
			fmt.Println("skipping malformed row:", row)
			continue
		}

		user := User{
			Id:    row[0],
			Email: row[1],
			Name:  row[2],
			Phone: row[3],
		}

		fmt.Printf("User: %+v\n", user)
	}
}
