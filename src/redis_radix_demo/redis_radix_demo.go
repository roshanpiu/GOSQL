package main

import (
    "github.com/fzzy/radix/redis"
    "log"
)

func main()  {
    client, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        log.Fatal(err)
    }
    
    r := client.Cmd("SET", "foo", 1234)
    if r.Err != nil {
        log.Fatal(err)
    }
    
    log.Printf("r.String(): %s\n", r.String())
    
    r = client.Cmd("GET", "foo")
    if r.Err != nil {
        log.Fatal(err)
    }
    
    log.Printf("r.String(): %s\n", r.String())
    
    i, _ := r.Int()
    log.Printf("i: %d\n", i)
    
    //running batch processes
    client.Append("SET", "name", "Roshan")
    client.Append("GET", "name")
    
    //prints OK
    r = client.GetReply()
    log.Printf("r.String(): %s\n", r.String())
    
    //prints Roshan
    r = client.GetReply()
    log.Printf("r.String(): %s\n", r.String())
    
    client.Cmd("SET", "first_name", "Roshan")
    client.Cmd("SET", "last_name", "Bates")
    
    //reading multiple value at once
    r = client.Cmd("MGET", "first_name", "last_name")
    
    //r.List() returns a string we need to convert the values to the what ever the type we want
    list, _ := r.List()
    for i, m := range list {
        log.Printf("value %d: %s\n", i+1, m)
    }
    
    
    
    
    
}