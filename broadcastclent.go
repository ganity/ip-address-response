package main

import (
    "fmt"
    "net"

)

func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}

func main() {
    ServerAddr,err := net.ResolveUDPAddr("udp","255.255.255.255:9345")
    CheckError(err)

    LocalAddr, err := net.ResolveUDPAddr("udp", ":9345")
    CheckError(err)

    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    CheckError(err)

    defer Conn.Close()
       msg := "get_internal()"
       buf := []byte(msg)
       _, err = Conn.Write(buf)
       CheckError(err)

       startlisten()

}



func startlisten() {
    ServerAddr,err := net.ResolveUDPAddr("udp",":9346")
    CheckError(err)

    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)
    defer ServerConn.Close()

    buf := make([]byte, 1024)

    for {
        n,addr,err := ServerConn.ReadFromUDP(buf)
        fmt.Println("========================================================")
        fmt.Println("Received:\n",string(buf[0:n]), "from:",addr)
        fmt.Println("========================================================")
        if err != nil {
            fmt.Println("Error: ",err)
        } 
    }
}