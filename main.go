package main

import (
	"bufio"
	"fmt"
	"os"
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
	for sc.Scan() {
		arr = append(arr, strings.ToLower(sc.Text()))
	}
	return arr, sc.Err()

}

func main() {
	arr, err := getString("./data.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(arr)
}
