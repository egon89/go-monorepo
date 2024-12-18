package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
)

type User struct {
	Id    string
	Email string
	Name  string
	Phone string
}

func main() {
	csvPath := "/tmp/large_users.csv"
	readAllFile("./resource/users_1k.csv", false)
	readLineByLine("./resource/users_10k.csv", false)
	readLineByLineGoRoutine(csvPath, true)
}

func readAllFile(path string, execute bool) {
	if !execute {
		return
	}
	fmt.Println("> reading all file")

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

func readLineByLine(path string, execute bool) {
	if !execute {
		return
	}
	fmt.Println("> reading line by line")

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

func readLineByLineGoRoutine(path string, execute bool) {
	if !execute {
		return
	}
	fmt.Println("> reading line by line with go routine")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	brokerUser := os.Getenv("RABBITMQ_USER")
	brokerPassword := os.Getenv("RABBITMQ_PASS")
	brokerHost := os.Getenv("RABBITMQ_HOST")
	brokerPort := os.Getenv("RABBITMQ_PORT")

	brokerConnection := fmt.Sprintf("amqp://%s:%s@%s:%s", brokerUser, brokerPassword, brokerHost, brokerPort)

	conn, err := amqp091.Dial(brokerConnection)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// buffered reader to read the file line by line
	reader := csv.NewReader(bufio.NewReader(file))

	header, err := reader.Read()
	if err != nil {
		panic(err)
	}
	fmt.Println("Headers:", header)

	lines := make(chan []string)
	results := make(chan User)

	var wg sync.WaitGroup

	numWorkers := 4
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(lines, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	go func() {
		for {
			row, err := reader.Read()
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				fmt.Println("Error reading row:", err)
				continue
			}
			lines <- row
		}
		close(lines)
	}()

	for user := range results {
		fmt.Printf("User: %+v\n", user)
	}
}

func worker(lines <-chan []string, results chan<- User, wg *sync.WaitGroup) {
	defer wg.Done()
	for row := range lines {
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

		results <- user
	}
}
