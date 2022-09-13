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
	log.SetLevel(log.DebugLevel)
    log.SetFormatter(&log.TextFormatter{
        PadLevelText: true,
        FullTimestamp: true,

    })
    
	log.SetOutput(os.Stdout)
}

func handleCon(c net.Conn) {
	cf := lib.OpenServerConfig("server.toml")
    log.Info("NEW CLIENT : " , c.RemoteAddr().String())

	for {

		rawData, err := bufio.NewReader(c).ReadString('\n')
		nw := bufio.NewWriter(c)

		if err != nil {

			//fmt.Println(err)
            log.Fatal("Failed to read from Network buffer with err => " , err)
			return

		}
		tmp := strings.TrimSpace(string(rawData))
		rawRequest := strings.TrimSpace(tmp)


        log.Info("NEW REQ : " , rawRequest)

		raw_res := lib.ParseRequest(rawRequest, cf)
		res := lib.NewResponse(raw_res)
		
        nw.Write(res)
		nw.Flush()

		if rawRequest == "+0" { //Not related to any spec; just for convenience
			log.Info("CLIENT QUIT : ", c.RemoteAddr().String())
			break
		}
        log.Debug("NEW UNKNOWN CMD: " , rawRequest)
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
		log.Info("SERVER LIVE " , ADDRESS)
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
