package main

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//log.SetOutput(file)
	logger := slog.New(slog.NewTextHandler(file, nil))

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Начало работы")
	flag := "exit"
	for scanner.Scan() {
		line := scanner.Text()
		if line == flag {
			fmt.Println("Завершение работы")
			return
		}
		//log.Println(line)
		logger.Info(line)
	}
}
