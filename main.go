package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getString(fileName string) ([]string, error) {
	arr := make([]string, 0)
	file, err := os.Open(fileName)
	if err != nil {
		return arr, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		arr = append(arr, strings.ToLower(sc.Text()))
	}
	return arr, sc.Err()

}

func main() {
	arr, err := getString("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	countWords := make(map[string]int)
	for _, word := range arr {
		countWords[word]++
	}
	//err = os.WriteFile("output.csv", []byte("слово.частота\n"), 0644)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = os.WriteFile("output.csv", countWords, 0644)
	// 2. Создаем или открываем файл для записи
	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// 3. Создаем CSV writer
	writer := csv.NewWriter(file)
	writer.Comma = '.'
	defer writer.Flush()

	// 4. Записываем заголовок (опционально)
	if err := writer.Write([]string{"Слово", "Частота"}); err != nil {
		log.Fatal(err)
	}

	// 5. Итерация по map и запись строк
	for k, v := range countWords {
		if err := writer.Write([]string{k, strconv.Itoa(v)}); err != nil {
			log.Fatal(err)
		}
	}
	for word, count := range countWords {
		fmt.Printf("%s.%d\n", word, count)
	}
	fmt.Println(arr)
	fmt.Println(len(arr))

}
