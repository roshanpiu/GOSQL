package main

import (
    "encoding/csv"
    "log"
    "os"
)

func main()  {
    f, err := os.Create("scrap2.csv")
    if err != nil {
        log.Fatal(err)
    }
    
    //writing one line at a time
    w := csv.NewWriter(f)
    w.Write([]string{"first", "last", "email"})
    w.Flush()
    w.Write([]string{"Mark", "Bates", "mark@example.com"})
    w.Flush()
    
    //writing multiple lines at a time
    w.WriteAll([][]string{
        []string{"Mark", "Bates", "mark@example.com"},
        []string{"Jane", "smith", "jane@example.com"},
        []string{"Roshan", "piumal", "roshan@example.com"},
    })

}