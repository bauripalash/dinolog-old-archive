package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type SitePost struct {
	Uid     string
	Title   string
	Summary string
}

const summaryMaxLen = 150
const titleMaxLen = 50

func (s *SitePost) ToFmtString() string {
	output := fmt.Sprintf("--> %s <--\n= %s =\n\n%s\n\n", s.Uid, s.Title, s.Summary)
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
	temp_title := ""
	got_title := false

	f, err := os.Open(fpath)
	if err != nil {
		log.Fatalf("Failed to read file post")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if isTitle(scanner.Text()) {
			temp_title = strings.TrimPrefix(scanner.Text(), "# ")
			got_title = true
			continue
		}

		if got_title {
			temp_summary += scanner.Text()
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(temp_title) > titleMaxLen {
		temp_title = temp_title[:titleMaxLen] + "..."
	}

	if len(temp_summary) > summaryMaxLen {
		temp_summary = temp_summary[:summaryMaxLen] + "..."
	}

	output_post.Title = temp_title
	output_post.Summary = temp_summary

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

	return posts

}
