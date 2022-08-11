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
	cf := lib.OpenConfig("server.ini")
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
		//log.New

		if err != nil {
			fmt.Println(err)
		}

		if request[0] == "+out" {
			log.Info("Client Quit Request : ", c.RemoteAddr().String())
			break
		}

		if request[0] == "+posts" {
			if len(request) == 2 {
				site_name := request[1]
				log.Info("REQ Site: ", site_name)
				//log.Info("Dummy example")
				if cf.CheckIfSiteExists(site_name) {
					//nw.Write([]byte(cf.GetSitePath(site_name) + "\n"))
					//nw.Write([]byte(strconv.FormatBool(cf.SitePathExists(site_name)) + "\n"))
					site_conf, _ := lib.GetSiteConfig(cf, site_name)
                    tempsite , _ := lib.GetSite(&site_conf)
                    
                    posts := tempsite.ReadPosts()

                                       
                    nw.Write([]byte(tempsite.Title + "\n"))
                    
                    nw.Write([]byte("~~~~~~~~\n\n=== POSTS ===\n\n"))

                    for _,post := range posts{
                        
                        nw.Write([]byte(post.ToFmtString()))
                        nw.Write([]byte("\n----------------\n"))
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
	//cf := lib.OpenConfig("server.ini")
	//fmt.Println(cf.CheckIfSiteExists("mangoman"))
	//fmt.Println(cf.GetSitePath("mangoman"))

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
