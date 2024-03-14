// test via tcp :6379
package main

import (
    "net"
    "fmt"
)

func main() {
    fmt.Println("1) Testing TCP server is running properly")
    conn, err := net.Dial("tcp", net.JoinHostPort("localhost", "6379"))
    if err != nil {
        fmt.Println("X Failed")
        return
    }
    fmt.Println("+ Passed")
    defer conn.Close()
}
