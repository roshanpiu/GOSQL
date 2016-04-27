package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "time"
)

type Todo struct {
    id          int
    subject     string
    desctiption string
    completed   bool
    created     time.Time
    updated     time.Time
}

func main()  {
    db, err := sql.Open("postgres", "postgres://postgres:007@localhost/test?sslmode=disable")
    
    if err != nil {
        log.Fatal(err)
    }
    
    now := time.Now()
    res, err := db.Exec("insert into todos (subject, description, completed, created, updated) values ($1, $2, $3, $4, $5)", "wash the cloths", "wash the cloths", false, now, now)
    res, err = db.Exec("insert into todos (subject, description, completed, created, updated) values ($1, $2, $3, $4, $5)", "wash the cloths", "wash the cloths", false, now, now)
    
    if err != nil {
        log.Fatal(err)
    }
    affected, _ := res.RowsAffected()
    log.Printf("Rows affected %d", affected)
    
    
    var subject string 
    
    //querying multiple rows
    rows, err := db.Query("select subject from todos")
    
    for rows.Next() {
        err = rows.Scan(&subject)
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("Subject is %s", subject)
    }
    
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
    
    //querying only one rows
    row := db.QueryRow("select subject from todos where id = $1", 1)
    row.Scan(&subject)
    log.Printf("The Frist subject is %s", subject)
    
    //querying all rows and columns from tables
    rows, err = db.Query("select * from todos")
    
    for rows.Next() {
        todo := Todo{}
        err = rows.Scan(&todo.id, &todo.subject, &todo.desctiption, &todo.completed, &todo.created, &todo.updated)
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("Subject is %s", todo.subject)
    }
    
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
}