package lib

import (
	"sort"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

//Get Posts ALL -> dd~mangoman~all

type PostFilter struct {
	ListFilter string
	Extra      string
}

type Request struct {
	Sitename string
	Filter   PostFilter
}

const sp = "~"

func splitSep(s string) []string {

	return strings.Split(s, sp)
}

func handleSingleSite(rawcmds []string) Request {

	sitename := rawcmds[0]
	command := rawcmds[1]

	if strings.HasPrefix(command, "L") || strings.HasPrefix(command, "l") {
		return Request{Sitename: sitename, Filter: PostFilter{ListFilter: "L", Extra: command[1:]}}
	} else if strings.HasPrefix(command, "ALL") || strings.HasPrefix(command, "all") {
		return Request{Sitename: sitename, Filter: PostFilter{ListFilter: "A", Extra: command[3:]}}

	} else if strings.HasPrefix(command, "O") || strings.HasPrefix(command, "o") {
		return Request{Sitename: sitename, Filter: PostFilter{ListFilter: "O", Extra: command[1:]}}

	} else if strings.HasPrefix(command, "T") || strings.HasPrefix(command, "t") {
		return Request{Sitename: sitename, Filter: PostFilter{ListFilter: "T", Extra: command[1:]}}

	} else if strings.HasPrefix(command, "D") || strings.HasPrefix(command, "d") {
		return Request{Sitename: sitename, Filter: PostFilter{ListFilter: "D", Extra: command[1:]}}

	}

	return Request{}
}

func parseDD(rawcmds []string, cfg ServerConfig) string {
	output := ""
	//log.Info("parseDD -> %s", rawcmds)
	if rawcmds[0] == "ST" || rawcmds[0] == "st" {
		req := handleSingleSite(rawcmds[1:])

		siteName := req.Sitename
		var siteconf SiteConfig
		var site Site

		if cfg.CheckIfSiteExists(siteName) {
			siteconf, _ = cfg.GetSiteConf(siteName)

			site = siteconf.GetSite()
		}

		posts := site.ReadPosts()
		limit := len(posts)

		if req.Filter.ListFilter != "D" {
			_, err := strconv.Atoi(req.Filter.Extra)
			if err != nil {
				log.Warn("Failed to pass argument as integer")
				limit = len(posts)
				//return ""
			}

		}

		switch filter := req.Filter.ListFilter; filter {

		case "A":

			//posts := site.ReadPosts()
			for _, post := range posts {

				output += post.ToFmtString()

			}

		case "O":

			//fmt.Println(">>>>>> ",req.Filter.Extra , " <<<<<<<<<")

			if limit >= len(posts) {
				limit = len(posts)
			}

			sort.SliceStable(posts, func(i, j int) bool {

				return posts[i].Date.Before(posts[j].Date)

			})

			for _, post := range posts[:limit] {
				output += post.ToFmtString()
			}

		case "L":
			for _, post := range posts[:limit] {
				output += post.ToFmtString()
			}

		case "D":
			arg_date, _ := time.Parse("2006-01-02", req.Filter.Extra)
			for _, post := range posts[:limit] {
				if post.Date.Format("2006-01-02") == arg_date.Format("2006-01-02") {
					output += post.ToFmtString()
				}

			}

		case "T":
			tags := strings.Split(req.Filter.Extra, ",")

			for _, post := range posts {
				for _, t := range tags {
					for _, pt := range post.Tags {
						if t == pt {
							output += post.ToFmtString()
						}
					}
				}
			}

		}

	} else if rawcmds[0] == "ID" || rawcmds[0] == "id" {
		println("Hey, I know the slug/id of the item")
	}

	return output

}

func ParseRequest(rawreq string, cfg ServerConfig) string {
	output := ""
	reqToken := splitSep(rawreq)

	if reqToken[0] == "dd" || reqToken[0] == "DD" {

		output = parseDD(reqToken[1:], cfg)

	}

	return output

}

func ReqDemo() {
	//	print("dd~s~mangoman~all")
}
