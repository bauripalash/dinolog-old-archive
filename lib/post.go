package lib

import (
	"fmt"
	"os"
	"path/filepath"
)


type SitePost struct{
    uid string
    title string
    summary string
    
}

func (s *Site) ReadPosts(){
    
    cdir := s.Contentdir //where all posts will be
    files, _ := os.ReadDir(cdir)

    cdir_path,_ := filepath.Abs(cdir) //absolute path to content dir 

    for _, file := range files{
        fmt.Println(filepath.Join(cdir_path, file.Name()))
    }

}
