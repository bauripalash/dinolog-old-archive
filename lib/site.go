package lib

import (
	"fmt"
	"os"
    log "github.com/sirupsen/logrus"
	"github.com/pelletier/go-toml/v2"
)

type SiteConfig struct {
	site_title  string
	sitepath    string
	content_dir string
}

type SiteConf struct {
	Site_name string
	Config    map[string]interface{}
}

type Site struct {
	cfg        SiteConfig
	Title      string
	Sitepath   string
	Contentdir string
}

func ReadSiteConfig(filepath string) SiteConf {

	var cfg SiteConf

	raw_conf, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Error Occured while opening config file %s", filepath)
	}

	toml.Unmarshal(raw_conf, &cfg)

	return cfg

}

func (s *SiteConf) GetConfigValue(target string) (string, bool) {

	target_value, ok := s.Config[target]

	if ok {
		return target_value.(string), true
	}

	return "", false

}

func (s *SiteConf) GetContentDir() (string, bool) {

	return s.GetConfigValue("content_dir")

}

func (s *SiteConf) GetSiteTitle() string {

	return s.Site_name

}

func (s *ServerConfig) GetSiteConf(sitename string) (SiteConfig, bool) {

	site_path, _ := s.GetSitePath(sitename)
	site_conf_file_path := site_path + "/site.toml"

	if doesConfigExist(site_conf_file_path) {
		conf_data := ReadSiteConfig(site_conf_file_path)
		fmt.Println(conf_data)
		content_dir, noerr_content_dir := conf_data.GetContentDir()
		site_title := conf_data.GetSiteTitle()

		if !noerr_content_dir {
			log.Fatalln("Content_Dir or Site_Title has errors")
			return SiteConfig{}, false
		}

		return SiteConfig{
			sitepath:    site_path,
			content_dir: content_dir,
			site_title:  site_title,
		}, true
	}

	log.Fatalln("Config Doesnot exists")
	return SiteConfig{}, false

}

func (sc *SiteConfig) GetSite() Site {

	return Site{
		Sitepath:   sc.sitepath,
		Title:      sc.site_title,
		Contentdir: sc.content_dir,
		cfg:        *sc,
	}

}

func (c *Site) GetContentDir() (string, bool) {
	c_dir := c.cfg.content_dir

	content_dir := c.cfg.sitepath + "/" + c_dir
	_, err := os.Stat(content_dir)

	//fmt.Println(err)
	if os.IsNotExist(err) {
		return "", false
	}
	return content_dir, true
}
