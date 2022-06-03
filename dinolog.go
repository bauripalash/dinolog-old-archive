package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type DlogEntry struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type Dlog struct {
	Name  string      `json:"name"`
	Host  string      `json:"host"`
	Uname string      `json:"uname"`
	Posts []DlogEntry `json:"entries"`
}

func (p *DlogEntry) formatEntry() string {
    var hasTitle string = "0"
    if len(p.Title) > 1 { hasTitle = "1" }
    return fmt.Sprintf("E %s\r\n\n%s\r\n\n%s\r\n\n" , hasTitle , p.Title, p.Text)

}

func (p *Dlog) formatDlog() string{
    
    var hasName string = "N"
    var hasHost string = "H"
    var hasUname string = "U"
    var hasPosts string = "P"
    if len(p.Name) > 1 { hasName+="1" }
    if len(p.Host) > 1 { hasHost+= "1" }

    if len(p.Uname) > 1 { hasUname+= "1" }

    if len(p.Posts) > 1 { hasPosts+= fmt.Sprintf("%d" , len(p.Posts)) }
    
    bStatus := fmt.Sprintf("D %d %d\r\n\n" , 0 , 1)
    bHeader := fmt.Sprintf("DLOG %s %s %s\r\n\n%s\r\n\n" , hasName , hasHost , hasUname , hasPosts)
    bReturn := fmt.Sprintf("%s%s" , bStatus , bHeader)
    
    //var bEntries []string

    for _,j := range p.Posts{
       //bEntries[i]=j.formatEntry()
       bReturn += j.formatEntry()
    }

    
    
    return bReturn

}

func getPosts() []byte {

	var myposts [2]DlogEntry

	myposts[0] = DlogEntry{
		Title: "Hello world",
		Text:  "Yuhuuu....",
	}

	myposts[1] = DlogEntry{
		Title: "my post",
		Text:  "Mew mew",
	}

	mylog := Dlog{
		Name:  "MyLog",
		Host:  "localhost",
		Uname: "palash",
		Posts: myposts[:],
	}
	fmt.Println(mylog)
	//myjson, err := json.Marshal(&mylog)
    
    
    //myjson := "\n[POST]\n"



	//fmt.Println("start->")
	//fmt.Println(err, string(myjson))
	return []byte(mylog.formatDlog())

}

func handleCon(c net.Conn) {

	fmt.Printf("hello from %s\n", c.RemoteAddr().String())

	for {

		rawData, err := bufio.NewReader(c).ReadString('\n')
		nw := bufio.NewWriter(c)

		if err != nil {

			fmt.Println(err)
			return

		}
		//fmt.Println(string(rawData))
		tmp := strings.TrimSpace(string(rawData))
        request := strings.Split(tmp , " ")
        
        fmt.Println(request)
        
        if err != nil {
			fmt.Println(err)
		}

		if tmp == "OUT" {
			break
		}

		if tmp == "POSTS" {
			nw.Write(getPosts())
			nw.Flush()
		}

		//res := fmt.Sprintf("%f", time.Hour.Minutes())

		c.Write([]byte(tmp))
	}
	c.Close()

}

func main() {

	x := true

	if x {
		var PORT int = 2001
		var ADDRESS string = fmt.Sprintf("127.0.0.1:%d", PORT)
		l, err := net.Listen("tcp4", ADDRESS)

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
