package main

import (
	"fmt"
	"net"
	"os"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("server port\n")
		return
	}

	sport := os.Args[1]
	ServerAddr, err := net.ResolveUDPAddr("udp", ":"+sport)
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		if n > 0 && n < 1024 {
			fmt.Println("Received ", string(buf[0:n-1]), " from ", addr)
			if err != nil {
				fmt.Println("Error: ", err)
			}

			buf2 := []byte(fmt.Sprintf("Your addr is: %v, your message is: %s", addr, string(buf[0:n-1])))
			ServerConn.WriteToUDP(buf2, addr)
			if err != nil {
				fmt.Println("Write Error: ", err)
			}
		}
	}
}
