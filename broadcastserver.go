package main

import (
    "log"
    "net"
    "os"
    "strings"
    "bytes"
)

func CheckError(err error) {
    if err  != nil {
        log.Println("Error: " , err)
        os.Exit(0)
    }
}

func main() {
    
    var logFilename string =  "log.log";
    logFile, err := os.OpenFile(logFilename, os.O_RDWR | os.O_CREATE, 0777)
    if err != nil {
        os.Exit(-1)
    }
    defer logFile.Close()
    logger := log.New(logFile,"\r\n", log.Ldate | log.Ltime | log.Lshortfile)
    
    
    ServerAddr,err := net.ResolveUDPAddr("udp",":9345")
    CheckError(err)

    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)
    defer ServerConn.Close()

    buf := make([]byte, 1024)

    for {
        n,addr,err := ServerConn.ReadFromUDP(buf)
        data := string(buf[0:n])
        logger.Println("========================================================")
        logger.Println("Received:\n",data, "from:",addr)
        logger.Println("========================================================")
        if err != nil {
            logger.Println("Error: ",err)
        } 
       
       ip := strings.Replace(addr.String(), ":9345", ":9346", -1)
       raddr, _ := net.ResolveUDPAddr("udp", ip);
        Conn, err := net.DialUDP("udp", nil, raddr)
        CheckError(err)

        defer Conn.Close()   
        
        _, err = Conn.Write([]byte(get_internal()))
        CheckError(err)

        
        // ip := get_internal()
        // ServerConn.WriteToUDP([]byte(ip), addr)
    }
}

func get_internal() string {
    var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
    
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        os.Stderr.WriteString("Oops:" + err.Error())
        os.Exit(1)
    }
    for _, a := range addrs {
        if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                buffer.WriteString(ipnet.IP.String() + "\n")
            }
        }
    }
    //os.Exit(0)
    return buffer.String()
}
