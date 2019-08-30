package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	_ "github.com/tealeg/xlsx"
)

var inFile = "E:\\GoProjects\\go-practice\\excel\\assessment.xlsx"

func main() {
	// 打开文件
	file, err := xlsx.OpenFile(inFile)
	checkErr(err)

	// 遍历sheet 页读取
	for _, sheet := range file.Sheets {
		fmt.Println("sheet name: ", sheet.Name)
		// 遍历行读取
		for _, row := range sheet.Rows {
			// 遍历每行的列读取
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%60s", text)
			}
			fmt.Print("\n")
		}
	}
	fmt.Println("\n\n import success")
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
