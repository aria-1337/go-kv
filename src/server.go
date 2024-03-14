package main

import (
    "net"
    "fmt"
    "encoding/json"
)

// TODO: Actual connection with username/db name etc dont just allow random blind connections

type command struct {
    Type string `json:"type"`
    Key string `json:"key"`
    Value string `json:"value"`
}

type response struct {
    Message string `json:"message"`
    Value interface{} `json:"value"`
}


func handleConnection(conn net.Conn, encoder *json.Encoder, decoder *json.Decoder, mem map[string]string) {
    defer conn.Close()
    c := command{}
    for {
        err := decoder.Decode(&c)
        if err != nil {
            fmt.Println(err)
            break
        }
        switch c.Type {
            case "echo":
                encoder.Encode(response{ Message: "OK", Value: "ECHO" })
            case "set":
                set(mem, encoder, c)
            case "get":
                get(mem, encoder, c)
        }
    }
}

func set(mem map[string]string, encoder *json.Encoder, c command) {
    mem[string(c.Key)] = string(c.Value)
    encoder.Encode(response{ Message: "OK", Value: "" })
}

func get(mem map[string]string, encoder *json.Encoder, c command) {
    encoder.Encode(response{ Message: "OK", Value: mem[string(c.Key)] })
}

func main() {
    mem := make(map[string]string)
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
        go handleConnection(conn, encoder, decoder, mem)
    }
}
