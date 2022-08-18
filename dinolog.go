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
	//cf := lib.OpenConfig("server.ini")
	cf := lib.OpenServerConfig("server.toml")
	fmt.Printf("hello from %s\n", c.RemoteAddr().String())

	for {

		rawData, err := bufio.NewReader(c).ReadString('\n')
		nw := bufio.NewWriter(c)

		if err != nil {

			fmt.Println(err)
			return

		}
		tmp := strings.TrimSpace(string(rawData))
		rawRequest := strings.TrimSpace(tmp)
		request := strings.Split(tmp, " ")

		if err != nil {
			fmt.Println(err)
		}

		//println(rawRequest)
		log.Info(rawRequest)
		nw.Write([]byte(lib.ParseRequest(rawRequest, cf)))
		nw.Flush()

		if request[0] == "+out" {
			log.Info("Client Quit Request : ", c.RemoteAddr().String())
			break
		}

		if request[0] == "+posts" {
			lib.ReqDemo()
			if len(request) == 2 {
				site_name := request[1]
				log.Info("REQ Site: ", site_name)
				if cf.CheckIfSiteExists(site_name) {
					site_config, noerr := cf.GetSiteConf(site_name)

					fmt.Println(site_config, noerr)
					tempsite := site_config.GetSite()

					posts := tempsite.ReadPosts()
					fmt.Println(len(posts))

					nw.Write([]byte(tempsite.Title + "\n"))

					//nw.Write([]byte("~~~~~~~~\n\n=== POSTS ===\n\n"))

					for _, post := range posts {

						nw.Write([]byte(post.ToFmtString()))
						//nw.Write([]byte("\n----------------\n"))
					}

				} else {

					log.Warn("Requested site not present in the server ", site_name)
					nw.WriteString("ERR! Site not found\n")
				}
				//	nw.Flush()
			} else {

				log.Warn("No site name provided")
				nw.WriteString("Please provide a site name\n")
			}
			//site_conf,_ := lib.GetSiteConfig("./mangoman/site.ini")
			//title,_ := site_conf.GetSiteTitle()
			//log.Info("Site title =>")
			//log.Info(title)
			nw.Flush()

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
