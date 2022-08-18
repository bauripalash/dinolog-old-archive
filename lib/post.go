package lib

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/pelletier/go-toml/v2"
	log "github.com/sirupsen/logrus"
)

type SitePost struct {
	Uid     string
	Title   string
	Summary string
	Date    time.Time
}

type SitePostMeta struct {
	Title string
	Date  string
}

const summaryMaxLen = 150
const titleMaxLen = 50
const ISODate = "2006-01-02T15:04:05-0700"

func (s *SitePost) ToFmtString() string {
	output := fmt.Sprintf("%s\n%s\n%s\n\n", s.Uid, s.Title, s.Summary)
	return output
}

func isTitle(t string) bool {
	text := strings.TrimSpace(t)
	if strings.HasPrefix(text, "# ") {
		return true
	}
	return false
}

func makePost(fpath string, uid string) SitePost {

	output_post := SitePost{
		Uid: uid,
	}

	temp_summary := ""
	temp_meta := ""

	f, err := os.Open(fpath)
	if err != nil {
		log.Fatalf("Failed to read file post")
	}

	defer f.Close()

	inside_meta := false

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if !inside_meta {

			if scanner.Text() == "++++" {

				inside_meta = true
				continue
			}

			if len(scanner.Text()) > 2 && len(temp_summary) < summaryMaxLen {
				temp_summary = scanner.Text()

			}

		} else if inside_meta {

			if scanner.Text() == "++++" {
				inside_meta = false
				continue
			}

			temp_meta += scanner.Text() + "\n"

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var metaData SitePostMeta

	toml.Unmarshal([]byte(temp_meta), &metaData)

	if len(metaData.Title) > 1 {
		output_post.Title = metaData.Title
	}

	if len(metaData.Date) > 1 {

		datetime, err := time.Parse(ISODate, metaData.Date)

		if err != nil {
			log.Fatalf(err.Error())
		}

		//log.Warn(datetime.Format("January 02, 2006"))
		output_post.Date = datetime
		output_post.Title += " / " + datetime.Format("January 02, 2006")

	}

	//logrus.Info(temp_meta)
	//logrus.Warn("Post Title -> " , metaData.Title)
	output_post.Summary = temp_summary

	//output_post.MetaData = metaData

	return output_post

}

func (s *Site) ReadPosts() []SitePost {
	posts := []SitePost{}
	cdir, noerr := s.GetContentDir()
	if !noerr {
		log.Fatalln("error getting content dir")
	}

	files, _ := os.ReadDir(cdir)

	cdir_path, _ := filepath.Abs(cdir) //absolute path to content dir

	for _, file := range files {
		posts = append(posts, makePost(filepath.Join(cdir_path, file.Name()), file.Name()))

	}
	sort.SliceStable(posts, func(i, j int) bool {

		return posts[i].Date.After(posts[j].Date) // Latest First

	})

	return posts

}
