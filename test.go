// test via tcp :6379
package main

import (
    "net"
    "fmt"
    "encoding/json"
    "os"
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

func main() {
    fmt.Println("1) Testing TCP server is running properly")
    conn, err := net.Dial("tcp", net.JoinHostPort("localhost", "6379"))
    if err != nil {
        fmt.Println("X Failed")
        os.Exit(1)
    }
    defer conn.Close()
    fmt.Println("+ Passed")

    encoder := json.NewEncoder(conn)
    decoder := json.NewDecoder(conn)

    fmt.Println("2) We are able to send an echo command and get the correct response")
    str := &command{
        Type: "echo",
        Key: "",
        Value: "test"}

    encoder.Encode(str)

    r := response{}
    echoErr := decoder.Decode(&r)
    if echoErr != nil {
        fmt.Println("X Failed", echoErr)
        os.Exit(1)
    }
    if r.Value != "ECHO" || r.Message != "OK" {
        fmt.Println("X Failed. Given: ", r.Value, "expected: ECHO ", "Message: ", r.Message)
        os.Exit(1)
    }
    fmt.Println("+ Passed")
}
