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

type Response struct{
    Status uint
    Size uint
    NumPost uint //Number of posts
    Data string

}

const SEPARATOR = "~" //Request separator
const DATEFMT = "2006-01-02"

func splitSep(s string) []string {

	return strings.Split(s, SEPARATOR)
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

func parseDD(rawcmds []string, cfg ServerConfig) (string,uint) {
	output := ""
    output_post_len := 0
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
            output_post_len = len(posts)

		case "O":

			//fmt.Println(">>>>>> ",req.Filter.Extra , " <<<<<<<<<")

			if limit >= len(posts) {
				limit = len(posts)
			}

			sort.SliceStable(posts, func(i, j int) bool {

				return posts[i].Date.Before(posts[j].Date)

			})

			for _, post := range posts[:limit] {
                output_post_len += 1
				output += post.ToFmtString()
			}

            //output_post_len = len(posts[:limit])

		case "L":
			for _, post := range posts[:limit] {
				output_post_len += 1 
                output += post.ToFmtString()
                
			}

           // output_post_len = len(posts[:limit])

		case "D":
			arg_date, _ := time.Parse(DATEFMT, req.Filter.Extra)
			for _, post := range posts[:limit] {
				if post.Date.Format(DATEFMT) == arg_date.Format(DATEFMT) {
                    output_post_len += 1
					output += post.ToFmtString()
				}

			}

		case "T":
			tags := strings.Split(req.Filter.Extra, ",")

			for _, post := range posts {
				for _, t := range tags {
					for _, pt := range post.Tags {
						if t == pt {
                            output_post_len += 1
							output += post.ToFmtString()
						}
					}
				}
			}

		}

	} else if rawcmds[0] == "ID" || rawcmds[0] == "id" {
        site_name := rawcmds[1]
        post_id := rawcmds[2]
        
        //site_path,_ := cfg.GetSitePath(site_name)
       // post_path := site_path + post_id

    	var siteconf SiteConfig
		var site Site


    	if cfg.CheckIfSiteExists(site_name) {
			siteconf, _ = cfg.GetSiteConf(site_name)

			site = siteconf.GetSite()
		}
        
        all_posts := site.ReadPosts()

        for _, post := range all_posts{
            
            if post.Uid == post_id{
                output = site.GetSinglePost(post).Text
                output_post_len = 1
            }
        }

        
        

       // all_posts := 


	}
    //println(output_post_len)
	return output,uint(output_post_len)

}

func ParseRequest(rawreq string, cfg ServerConfig) Response {
	output := ""
    statuscode := 0
    number_of_posts := uint(0)
	reqToken := splitSep(rawreq)

	if reqToken[0] == "dd" || reqToken[0] == "DD" {

		output,number_of_posts = parseDD(reqToken[1:], cfg)

	}
    
    if len(output) > 1{
        statuscode = 1
    }
    // DD ~ STATUSCODE ~ SIZE ~ NUMBER OF POSTS ~ X ~ NUMBER OF OPTIONS
    //response := fmt.Sprintf("D~%d~%d~%d\r\n%s" , statuscode , len(output) , number_of_posts , output)

    //println(len(output))
	return Response{
        Status: uint(statuscode),
        NumPost: number_of_posts,
        Size: uint(len(output)),
        Data: output,
    }

}

