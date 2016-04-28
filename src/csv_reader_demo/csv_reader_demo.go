package main

import (
    "os"
    "log"
    "encoding/csv"
)

func main()  {
    f, err := os.Open("scrap.csv")
    if err != nil {
        log.Fatal(err)
       
    }
    
    r := csv.NewReader(f)
    r.Read() //throw away the headers
    
    for {
        row, err := r.Read()
        if err != nil {
            log.Println(err)
            break
        }
        printRow(row)
    }
    
    //reading all the rows at once
    recs, err := r.ReadAll()
    if err != nil {
        log.Fatal(err)
    }
    
    for _, row := range recs {
        printRow(row)
    }
    
    
    
}

func printRow(row []string) {
    log.Printf("len(row) %d\n", len(row))
    for i, col := range row {
        log.Printf("[%d]: %s\n", i , col)
    }
}