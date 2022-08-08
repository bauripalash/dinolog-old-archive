package lib

import (
	"log"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

type Config struct {
	pathname     string
	configobject ini.File
}

func doesConfigExist(f string) bool {
	fileinfo, err := os.Stat(f)

	if os.IsNotExist(err) {
		log.Fatalf("Config file can not be found!")
	} else if fileinfo.Size() == 0 {
		log.Fatalf("Empty config file")
	}
	return true

}

func OpenConfig(filename string) Config {

	doesConfigExist(filename)

	rawconfig, err := ini.Load(filename)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return Config{
		pathname:     filename,
		configobject: *rawconfig,
	}

}

func (c *Config) CheckIfSiteExists(sitename string) bool {
	if !c.configobject.HasSection("configuration") {
		log.Fatalf("Configuration section does not exist in config file")
	}

	if !c.configobject.Section("configuration").HasKey("available_sites") {
		log.Fatalf("Can not find available_sites configuration")
	}

	sitesraw := c.configobject.Section("configuration").Key("available_sites").Value()
	sites := strings.Split(sitesraw, ",")

	for _, s := range sites {
		if strings.TrimSpace(s) == sitename {
			return true
		}
	}
	return false
}

func (c *Config) GetSitePath(sitename string) string {
	//fmt.Println(sitename)
	if c.configobject.HasSection(sitename) {
		if c.configobject.Section(sitename).HasKey("root") {
			return c.configobject.Section(sitename).Key("root").Value()
		} else {
			log.Fatalf("Site configuration has no root specified")
		}
	} else {
		log.Fatalf("Config has no configuration for that site")
	}
	return ""
}

func (c *Config) SitePathExists(sitename string) bool {
	sitepath := c.GetSitePath(sitename)
	_, err := os.Stat(sitepath)

	if os.IsNotExist(err) {
		return false
	}
	return true
}
