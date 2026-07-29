package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"pack/modules"

	"github.com/xuri/excelize/v2"
)

func main() {
	data, err := os.ReadFile("../test.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &modules.Employees)
	if err != nil {
		log.Fatal(err)
	}

	f := excelize.NewFile()

	_ = f.SetCellValue("Sheet1", "A1", "№")
	_ = f.SetCellValue("Sheet1", "B1", "Имя")
	_ = f.SetCellValue("Sheet1", "C1", "Возраст")
	_ = f.SetCellValue("Sheet1", "D1", "Должность")
	_ = f.SetCellValue("Sheet1", "E1", "Зарплата")
	_ = f.SetCellValue("Sheet1", "F1", "Почта")
	_ = f.SetCellValue("Sheet1", "G1", "Статус")

	for i, v := range modules.Employees {
		row := i + 2
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), v.ID)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), v.Name)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), v.Age)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), v.Position)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), v.Salary)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), v.Email)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("G%d", row), v.IsActive)
	}
	
	err = f.SaveAs("report.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Файл создан")

}
