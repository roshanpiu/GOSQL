package main

import (
    "log"
    "github.com/boltdb/bolt"
    "os"
    "fmt"
    "encoding/json"
    "strconv"
)

//boltdb supports transactions

var (
    db_name     =   []byte("metacast")
    db              *bolt.DB
    bucketName  =   []byte("bucketName")
)

func init()  {
    var err error 
    db, err = bolt.Open("bolt.db", 0644, nil)
    if err != nil {
        log.Fatal(err)
    }
}

type Person struct {
    Name  string
    Email string  
}

func main()  {
    defer db.Close()
    defer os.Remove(db.Path())
    
    err := db.Update(func (tx *bolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists(bucketName)
        return err
    })
    
    if err != nil {
        log.Fatal(err)
    }
    
    err = db.Update(func (tx *bolt.Tx) error {
        b := tx.Bucket(bucketName)
        p := Person{"Mark Bates", "mark@example.com"}
        by, _ := json.Marshal(p)
        
        return b.Put([]byte("mark"),by)
        
    })
    
    db.View(func (tx *bolt.Tx) error {
        b := tx.Bucket(bucketName)
        by := b.Get([]byte("mark"))
        p := Person{}
        json.Unmarshal(by, &p)
        
        fmt.Printf("p: %s\n", p)
        return nil
    })
    
    //defining our own transactions
    // tx , _ := db.Begin(false)
    // b := tx.Bucket(bucketName)
    // by := b.Get([]byte("mark"))
    // p := Person{}
    // json.Unmarshal(by, &p)
    
    // fmt.Printf("p: %s\n", p)
    // tx.Commit()

    db.Update(func (tx *bolt.Tx) error {
        b := tx.Bucket(bucketName)
        for i:=0; i<10; i++ {
            b.Put([]byte(fmt.Sprintf("key-%d", i)), []byte(strconv.Itoa(i)))
        }
        return nil
    })
    
    db.View(func (tx *bolt.Tx) error {
        b := tx.Bucket(bucketName)
        
        //using foreach to iterate through a collection
        b.ForEach(func (k []byte, v []byte) error {
            fmt.Printf("%s = %s\n", k, v)
            return nil
        })
        
        //using cursors to iterate through a collection
        c := b.Cursor()
        for k, v := c.First(); k != nil; k, v = c.Next(){
            fmt.Printf("%s = %s\n", k, v)
        }
        
        //printing the first item
        k, v := c.First()
        fmt.Printf("%s = %s\n", k, v)
        
        //printing the first item
        k, v = c.Last()
        fmt.Printf("%s = %s\n", k, v)
        
        //prefix scans can be used to check whether a key starts with a specific prefix
        //range scans for searching for time ranges
        
        
        return nil
        
    })
    
    
    
}