package lib

import (
	"fmt"
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

/*
var STATUS_CODES map[string]int
STATUS_CODES["OK"] = 10
STATUS_CODES["SERVERERROR"] = 11
STATUS_CODES["MOVED"] = 12
STATUS_CODES["NOTFOUND"] = 13
*/

func (p *DlogEntry) FormatEntry() string {
	return fmt.Sprintf("E~%d~%s\r\n\n%s\r\n\n%s\r\n\n", p.Size, p.Slug, p.Title, p.Text)

}

func (l *Dlog) InsertNewEntry(entry *DlogEntry) {
	l.Posts = append(l.Posts, *entry)
}

func (p *Dlog) FormatDlog() string {

	bStatus := fmt.Sprintf("D~%d~%d\r\n", 10, len(p.Posts))
	bStatus += fmt.Sprintf("name~%s\r\n", p.Name)
	bStatus += fmt.Sprintf("uname~%s\r\n", p.Uname)
	bReturn := fmt.Sprintf("%s\r\n", bStatus)

	for _, j := range p.Posts {
		bReturn += j.FormatEntry()
	}

	return bReturn

}
