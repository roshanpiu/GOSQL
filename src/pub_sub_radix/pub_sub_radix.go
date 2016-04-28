package main

import (
    "github.com/fzzy/radix/redis"
    "github.com/fzzy/radix/extra/pubsub"
    "log"
    "fmt"
    "time"
)

func main()  {
    go func ()  {
        client, _ := redis.Dial("tcp", "localhost:6379")
        i := 0
        for {
            i++
            client.Cmd("PUBLISH", "news.tech", fmt.Sprintf("This is tech story #%d", i))
            client.Cmd("PUBLISH", "news.tech", fmt.Sprintf("This is tech story #%d", i))
        }
    }()
    
    go func ()  {
        client, _ := redis.Dial("tcp", "localhost:6379")
        sub := pubsub.NewSubClient(client)
        sr := sub.PSubscribe("news.*")
        
        if sr.Err != nil {
            log.Fatal(sr.Err)
        }
        
        for {
            r := sub.Receive()
            if r.Err != nil {
                log.Fatal(r.Err)
            }
            log.Printf("r.Message: %s\n", r.Message)
        }
    }()
    
    time.Sleep(1 * time.Second)
}