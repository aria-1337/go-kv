package main

import (
    "net"
    "fmt"
    "encoding/json"
)

type command struct {
    Type string `json:"type"`
    Key string `json:"key"`
    Value string `json:"value"`
}

type response struct {
    Message string `json:"message"`
    Value string `json:"value"`
}

func handleConnection(conn net.Conn, encoder *json.Encoder, decoder *json.Decoder) {
    defer conn.Close()
    c := command{}
    for {
        err := decoder.Decode(&c)
        if err != nil {
            fmt.Println(err)
            break
        }
        fmt.Println(c.Type)
    }
}


func main() {
    ln, err := net.Listen("tcp", ":6379")
    if err != nil {
        fmt.Println(err)
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
        }

        decoder := json.NewDecoder(conn)
        encoder := json.NewEncoder(conn)
        go handleConnection(conn, encoder, decoder)
    }
}
