package main

import (
    "os"
    "fmt"
    "github.com/tealeg/xlsx"
    "strconv"
    "time"
)

func main() {
    fmt.Println("[" + time.Now().Format("15:04:05") + "] Starting..")
    xlsx.PrintMemUsage()
    rowsN := os.Args[1]
    n, err := strconv.Atoi(rowsN)
    if err != nil {
        fmt.Println("Wrong argument.")
        return
    }

    sheetName := "Sheet1"

    file := xlsx.NewFile()
    _, err = file.AddSheet(sheetName)
    if err != nil {
        fmt.Printf(err.Error())
    }

    fmt.Println("[" + time.Now().Format("15:04:05") + "] Starting to generate document..")
    xlsx.PrintMemUsage()
    for i:=0; i < n; i++ {
        row := file.Sheet[sheetName].AddRow()

        for j:=0; j< 16; j++ {
            cell := row.AddCell()
            cell.Value = strconv.Itoa(i)
        }
    }

    fmt.Println("[" + time.Now().Format("15:04:05") + "] Document generated.")
    xlsx.PrintMemUsage()
    fmt.Println("[" + time.Now().Format("15:04:05") + "] Starting to generate file..")

    err = file.Save("MyXLSXFile.xlsx")

    fmt.Println("[" + time.Now().Format("15:04:05") + "] File generated.")
    if err != nil {
        fmt.Printf(err.Error())
    }
}