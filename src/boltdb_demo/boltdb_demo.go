package main

import (
    "log"
    "github.com/boltdb/bolt"
    "os"
    "errors"
    "fmt"
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

func main()  {
    defer db.Close()
    defer os.Remove(db.Path())
    
    //when use db.UWpdate db is in a writable state
    err := db.Update(func (tx *bolt.Tx) error {
        //every thing inside this function will be wrapped around a transaction
        //bolt will give us a transaction and expects us to return a error 
        b, err := tx.CreateBucketIfNotExists(bucketName)
        if err != nil {
            return err
        }
        //everything with bolt is a byte array 
        //everything we input to bolt we should wrap it around with a byte array
        return b.Put([]byte("mark"), []byte("mark bates"))
    })
    
    if err != nil {
        log.Fatal(err)
    }
    
    err = db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket(bucketName)
        b.Put([]byte("mark"), []byte("mark bates"))
        
        return errors.New("oops!!")
    })
    
    //When use db.View db is in 
    err = db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket(bucketName)
        m := b.Get([]byte("mark"))
        fmt.Printf("m: %s\n", m)
        return nil
    })
}