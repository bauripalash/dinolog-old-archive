package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type DlogEntry struct {
	Slug  string `json:"slug"`  //Unique Identifer of Log entry
	Title string `json:"title"` //Title (Optional)
	Text  string `json:"text"`  //Body
	Size  uint64 `json:"size"`  //Size of Text Body
}

type Dlog struct {
	Name  string      `json:"name"`
	Uname string      `json:"uname"`
	Posts []DlogEntry `json:"entries"`
}

func NewEntry(title, text, slug string) DlogEntry {
	return DlogEntry{
		Slug:  slug,
		Title: title,
		Text:  text,
		Size:  uint64(len(text)),
	}

}

func (p *DlogEntry) formatEntry() string {
	var hasTitle string = "0"
	if len(p.Title) > 1 {
		hasTitle = "1"
	}
	return fmt.Sprintf("E %s\r\n\n%s\r\n\n%s\r\n\n", hasTitle, p.Title, p.Text)

}

func (l *Dlog) InsertNewEntry(entry *DlogEntry) {
	l.Posts = append(l.Posts, *entry)
}

func (p *Dlog) formatDlog() string {

	bStatus := fmt.Sprintf("D~%d~%d\r\n", 01, len(p.Posts))
	bStyled := "=======MANGO========"
	bReturn := fmt.Sprintf("%s\r\n", bStatus)
	bReturn += bStyled
	//var bEntries []string

	for _, j := range p.Posts {
		//bEntries[i]=j.formatEntry()
		bReturn += j.formatEntry()
	}

	return bReturn

}

func getPosts() []byte {

	var myposts [2]DlogEntry

	myposts[0] = NewEntry("Hello world", "my first post", "1")

	myposts[1] = NewEntry("Bye world", "my last post", "2")

	mylog := Dlog{
		Name:  "MANGO",
		Uname: "palash",
	}
	mylog.InsertNewEntry(&myposts[0])
	mylog.InsertNewEntry(&myposts[1])
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
		request := strings.Split(tmp, " ")

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
