/*
History

Config

Encryption

Accounts


Multiple chats on single server

Text formating

File transfer
*/


package main


import (
	"fmt"
	"net"
	"bufio"
)


const (
	servAddr = ":44821"

    reset = "\033[0m"
    bold = "\033[1m"
    underline = "\033[4m"
    strike = "\033[9m"
    italic = "\033[3m"
	blinking = "\033[25m"

    cRed = "\033[31m"
    cGreen = "\033[32m"
    cYellow = "\033[93m"
    cBlue = "\033[34m"
    cPurple = "\033[35m"
    cCyan = "\033[36m"
    cWhite = "\033[37m"
)


var conns []net.Conn


func main() {
	l, err := net.Listen("tcp", servAddr)
	p(err)
	defer l.Close()

	//append(conns, os.Stdout)

	for {
		conn, err := l.Accept()
		p(err)

		defer conn.Close()

		conns = append(conns, conn)

		go handle(conn)
	}
}


func handle(conn net.Conn) {
	addr := conn.RemoteAddr().String()
	Broadcast(addr, true, "connected")
	reader := bufio.NewReader(conn)

	//m, _, err := reader.ReadLine()
	//if err != nil {panic(err)}

	for {
		m, _, err := reader.ReadLine()
		if err != nil {break}

		Broadcast(addr, false, string(m))
	}
	RemoveConn(conn)
	Broadcast(addr, true, "disconnected")
}


func Broadcast(who string, isInfo bool, msg string) {
	m := FmtMsg(who, isInfo, msg)
	fmt.Print(m)
	for _, conn := range conns {
		_, err := conn.Write([]byte(m))
		p(err)
	}
}


func FmtMsg (who string, isInfo bool, msg string) string {
	var s string
	if isInfo {
		s = italic + cYellow + " # "
	} else {
		s = " @ "
	}

	return bold + cGreen + who + reset + s + msg + reset + "\n"
}


func RemoveConn(conn net.Conn) {
	for i, curConn := range conns {
		if conn == curConn {
			conns = append(conns[:i], conns[i+1:]...)
		}
	}
}


func p(err error) {
	if err != nil {panic(err)}
}
