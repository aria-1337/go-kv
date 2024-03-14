package main

import (
    "net"
    "fmt"
)

func handleConnection(conn net.Conn) {
    defer conn.Close()
    buf := make([]byte, 4096)
    for {
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Println(err)
        }
        conn.Write(buf[:n])
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
        go handleConnection(conn)
    }
}
