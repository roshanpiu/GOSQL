package main

import (
    "log"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
    "database/sql"
    "encoding/json"
    "os"
    "database/sql/driver"
    "strings"
    
)

type User struct {
    ID      int                 `db:"id" json:"id"`
    Email   string              `db:"email" json:"email"`
    Name    MyNullString        `db:"name" json:"name"`
    
}

//creating a new NullString type inorder to handle the null strings

type MyNullString sql.NullString

func (ns *MyNullString) Scan(value interface{}) error{
    n := sql.NullString{String: ns.String}
    err := n.Scan(value)
    ns.String, ns.Valid = n.String, n.Valid
    return err
}

func (ns *MyNullString) Value() (driver.Value, error){
    n := sql.NullString{String: ns.String}
    return n.Value()
}

func (ns MyNullString) MarshalJSON() ([]byte, error) {
    if ns.Valid {
        return json.Marshal(ns.String)
    }
    return json.Marshal(nil)
}

func (ns *MyNullString) UnmarshalJSON(text []byte) error{
    ns.Valid = false 
    if string(text) == "null" {
        return nil
    }
    s := ""
    err := json.Unmarshal(text, &s)
    if err == nil {
        ns.Valid = true
        ns.String = s
    }
    return err
}

func main()  {
    
    db, err := sqlx.Open("postgres", "postgres://postgres:007@localhost/test?sslmode=disable")
    
    if err != nil {
        log.Fatal(err)
    }
    
    user := &User{}
    err = db.Get(user, "select * from users")
    
    if err != nil {
        log.Fatal(err)
    }
    
    log.Println("printing data retrived from the db")
    log.Println(user)
    log.Println("User name: ",user.Email)
    log.Println("User name: ",user.Name.String)
    log.Println("User name: ",user.Name.Valid)
    log.Printf("\n\n")
    json.NewEncoder(os.Stdout).Encode(user)
    log.Printf("\n\n")
    
    //unmarshaling json with the out new type
    log.Println("printing unmarhsaled json")
    u := User{}
    x := `{"id":1,"email":"ropiumal@gmail.com","name":"roshan"}`
    json.NewDecoder(strings.NewReader(x)).Decode(&u)
    log.Println(u)
    log.Println("User name: ", u.Email)
    log.Println("User name: ", u.Name.String)
    log.Println("User name: ", u.Name.Valid)
    
}