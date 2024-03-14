// test via tcp :6379
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

func main() {
    fmt.Println("1) Testing TCP server is running properly")
    conn, err := net.Dial("tcp", net.JoinHostPort("localhost", "6379"))
    if err != nil {
        fmt.Println("X Failed")
        return
    }
    defer conn.Close()
    fmt.Println("+ Passed")

    encoder := json.NewEncoder(conn)

    str := &command{
        Type: "echo",
        Key: "",
        Value: "test"}

    encoder.Encode(str)
}
