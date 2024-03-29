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
	Tags    []string
	Date    time.Time
	Path    string
}

type SitePostMeta struct {
	Title string
	Date  string
	Tags  []string
}

type SinglePost struct {
	Metadata SitePost
	Text     string
}

const SUMMARY_MAX_LEN = 150
const TITLE_MAX_LEN = 50
const ISO_DATE = "2006-01-02T15:04:05-0700"

func (s *SitePost) ToFmtString() string {
	output := fmt.Sprintf("%s\r\n%s\r\n%s\r\n%s\r\n\n", s.Uid, s.Title, s.Date.Format(DATEFMT) , s.Summary)
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

			if len(scanner.Text()) > 2 && len(temp_summary) < SUMMARY_MAX_LEN {
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

		datetime, err := time.Parse(ISO_DATE, metaData.Date)

		if err != nil {
			log.Fatalf(err.Error())
		}

		//log.Warn(datetime.Format("January 02, 2006"))
		output_post.Date = datetime
		//output_post.Title += " / " + datetime.Format("January 02, 2006")

	}

	if len(metaData.Tags) > 0 {
		output_post.Tags = metaData.Tags
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

func (s *Site) GetSinglePost(post SitePost) SinglePost {

	fulltext := ""

	cdir, _ := s.GetContentDir()
	f, err := os.Open(cdir + "/" + post.Uid)
	if err != nil {
		log.Fatalf("Failed to read file post")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	inside_meta := false
	for scanner.Scan() {

		if !inside_meta {
			if scanner.Text() == "++++" {
				inside_meta = true
				continue
			}

			fulltext += scanner.Text() + "\n"

		} else {
			if scanner.Text() == "++++" {

				inside_meta = false
				continue
			}
		}

	}

	return SinglePost{
		Metadata: post,
		Text:     fulltext,
	}

}
