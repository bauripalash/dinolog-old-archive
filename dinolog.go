package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/bauripalash/dinolog/lib"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetOutput(os.Stdout)
}

func handleCon(c net.Conn) {
	cf := lib.OpenServerConfig("server.toml")
	log.Info(fmt.Sprintf("NEW CLIENT : %s", c.RemoteAddr().String()))

	for {

		rawData, err := bufio.NewReader(c).ReadString('\n')
		nw := bufio.NewWriter(c)

		if err != nil {

			fmt.Println(err)
			return

		}
		tmp := strings.TrimSpace(string(rawData))
		rawRequest := strings.TrimSpace(tmp)

		if err != nil {
			fmt.Println(err)
		}

		log.Info(fmt.Sprintf("REQUEST : %s", rawRequest))
        raw_res := lib.ParseRequest(rawRequest, cf)
        res := lib.NewResponse(raw_res, true)
        nw.Write(res)
		nw.Flush()

		if rawRequest == "00" { //Not related to any spec; just for convenience
			log.Info("CLIENT QUIT : ", c.RemoteAddr().String())
			break
		}
		nw.WriteString("Unknown command\n")
	}
	c.Close()

}

func main() {
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
