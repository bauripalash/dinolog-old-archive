package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/bauripalash/dinolog/lib"
)

func handleCon(c net.Conn) {

	fmt.Printf("hello from %s\n", c.RemoteAddr().String())

	for {

		rawData, err := bufio.NewReader(c).ReadString('\n')
		nw := bufio.NewWriter(c)

		if err != nil {

			fmt.Println(err)
			return

		}
		tmp := strings.TrimSpace(string(rawData))
		request := strings.Split(tmp, " ")

		fmt.Println(request)

		if err != nil {
			fmt.Println(err)
		}

		if request[0] == "+out" {
			break
		}

		if request[0] == "+posts" {
			nw.Write(lib.GetPosts())
			nw.Flush()
		}
		c.Write([]byte(tmp))
	}
	c.Close()

}

func main() {
	lib.CreateDatabase()
	x := true

	if x {
		var PORT int = 2001
		var ADDRESS string = fmt.Sprintf("127.0.0.1:%d", PORT)
		l, err := net.Listen("tcp4", ADDRESS)
		fmt.Println("Starting server on 127.0.0.1:2001")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer l.Close()

		for {

			c, err := l.Accept()
			if err != nil {
				fmt.Println(err)
				return
			}
			go handleCon(c)

		}
	}

}
