package main

import (
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
    "log"
    "time"
)

type Todo struct {
    ID          int
    Subject     string
    Description string
    Completed   bool
    CreatedAt     time.Time `db:"created"`
    UpdatedAt     time.Time `db:"updated"`
}

func main()  {
    db, err := sqlx.Open("postgres", "postgres://postgres:postgres@localhost/test?sslmode=disable")
    
    if err != nil {
        log.Fatal(err)
    }
    
    //Transactions
    tx := db.MustBegin()
    now := time.Now()
    t := Todo{
        Subject:        "Mow Lawn",
        Description:    "Yuck!",
        Completed:       false,
        CreatedAt:      now,
        UpdatedAt:      now,
    }
    tx.MustExec("insert into todos (subject, description, completed, created, updated) values ($1, $2, $3, $4, $5)", t.Subject, t.Description, t.Completed, t.CreatedAt, t.UpdatedAt)
    tx.NamedExec("INSERT INTO todos (subject, description, completed, created, updated) values (:subject, :description, :completed, :created, :updated)", &t)
    tx.Commit()
    
    todos := []Todo{}
    err = db.Select(&todos, "select * from todos")
    if err != nil {
        log.Fatal(err)
    }
    
    for _, todo := range todos {
        log.Printf("Subject is %s Created is %s ", todo.Subject, todo.CreatedAt)
    }
}