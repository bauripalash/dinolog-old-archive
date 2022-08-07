package lib

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type SiteConfig struct{
    pathname string;
    sitepath string;
    configobject ini.File;
}

func GetSiteConfig(cfg Config , sitename string) (SiteConfig, error) {
    site_config_file := cfg.GetSitePath(sitename) + "/site.ini"    
fmt.Println(sitename)
   if doesConfigExist(site_config_file){
       raw_config := OpenConfig(site_config_file)
        return SiteConfig{
            pathname: raw_config.pathname ,
            configobject: raw_config.configobject,
            sitepath: cfg.GetSitePath(sitename) ,
        },nil 
   }else{
       return SiteConfig{} , fmt.Errorf("Config file %s cannot be found!" , site_config_file)
   }
}

func (c *SiteConfig) configHasContentDir() (string, bool){
    
    if c.configobject.HasSection("config"){
        if c.configobject.Section("config").HasKey("content_dir"){
            return c.configobject.Section("config").Key("content_dir").Value(),true
        }
    }
    return "",false

}

func (c *SiteConfig) getContentDir() (string,bool){
    
    content_dir , found := c.configHasContentDir()

    if found{
        return content_dir,true
    }else{
        log.Fatalf(fmt.Sprint("Site config doesn't have content directory mentioned."))
        return "",false
    }

}


func (c *SiteConfig) GetSiteContentDir() (string,bool){
    c_dir,nofail := c.getContentDir()
    fmt.Println(nofail)
    if !nofail{
        return "",false
    }
    
    content_dir := c.sitepath + "/" + c_dir 
    _ , err := os.Stat(   content_dir)
    
    fmt.Println(err)
    if os.IsNotExist(err){
        return "",false
    }
    return content_dir,true
}



func (c *SiteConfig) GetSiteTitle() (string , bool){
    //fmt.Println(c.configobject)
    if c.configobject.Section("").HasKey("site_name"){
        return c.configobject.Section("").Key("site_name").Value(),true
    }
    return "",false
    
}
