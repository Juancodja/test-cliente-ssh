package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("🍣 Bienvenido al proyecto Sushi!")

	conn, err := net.Dial("tcp", "localhost:2222")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	version := "SSH-2.0-SUSHI"
	fmt.Println("CLIENTE: ", version)

	fmt.Fprint(conn, version+"\r\n")

	msg, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("SERVIDOR:", msg)
}
