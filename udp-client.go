package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func localIP() (net.IP, error) {
	tt, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, t := range tt {
		aa, err := t.Addrs()
		if err != nil {
			return nil, err
		}
		for _, a := range aa {
			ipnet, ok := a.(*net.IPNet)
			if !ok {
				continue
			}
			v4 := ipnet.IP.To4()
			if v4 == nil || v4[0] == 127 { // loopback address
				continue
			}
			return v4, nil
		}
	}
	return nil, errors.New("cannot find local IP address")
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func innerTest() {
	local_ip := ""
	ip_port := ""
	if len(os.Args) == 1 {
		ip, err := localIP()
		if err == nil {
			local_ip = fmt.Sprintf("%v", ip)
		}
		ip_port = "sg.mjy211.com:12700"
		rip, _ := net.LookupHost("sg.mjy211.com")
		if rip != nil && len(rip) > 0 {
			ip_port = rip[0] + ":12700"
			fmt.Printf("resolv: %v\n", rip[0])
		}
	} else if len(os.Args) == 2 {
		ip, err := localIP()
		if err == nil {
			local_ip = fmt.Sprintf("%v", ip)
		}
		ip_port = os.Args[1]
	} else if len(os.Args) == 3 {
		ip_port = os.Args[1]
		local_ip = os.Args[2]
	} else {
		fmt.Printf("client remote_ip:port [localip]\n")
		return
	}
	fmt.Printf("Program by Xiaobao, using remote: %v\n", ip_port)
	rand.Seed(time.Now().UTC().UnixNano())

	//fmt.Print("Enter text: ")

	ServerAddr, err := net.ResolveUDPAddr("udp", ip_port)
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", local_ip+":0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	recvBuf := make([]byte, 1024)
	defer Conn.Close()

	//msg, _ := reader.ReadString('\n')
	msg := "random test " + randSeq(10)
	buf := []byte(msg)
	_, err = Conn.Write(buf)
	if err != nil {
		fmt.Println(msg, err)
		return
	}

	n, addr, err := Conn.ReadFromUDP(recvBuf)
	fmt.Println("Recv: ", string(recvBuf[0:n]), " from ", addr)
}

func main() {
	innerTest()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press ENTER to exit")
	reader.ReadString('\n')
}
