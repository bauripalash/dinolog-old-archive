package lib

import (
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type ServerConfig struct {
	Server_name   string
	Enabled_sites []string
	SiteConfig    []map[string]interface{}
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

func OpenServerConfig(filename string) ServerConfig {
	var cfg ServerConfig

	raw_conf, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error Occured while opening config file %s", filename)
	}
	//confdata := string(raw_conf)

	toml.Unmarshal(raw_conf, &cfg)

	return cfg

}

func (s *ServerConfig) IsSiteEnabled(sitename string) bool {
	for _, item := range s.Enabled_sites {
		if item == sitename {
			return true
		}
	}

	return false
}

func (s *ServerConfig) CheckIfSiteExists(sitename string) bool {

	for _, item := range s.SiteConfig {
		if item["name"] == sitename {
			return true
		}
	}

	return false

}

func (s *ServerConfig) GetSitePath(sitename string) (string, bool) {

	for _, site := range s.SiteConfig {
		if site["name"] == sitename {
			//return site
			return site["root"].(string), true
		}
	}
	return "", false
}

func (s *ServerConfig) SitePathExists(sitename string) bool {
	sitepath, no_err := s.GetSitePath(sitename)

	if !no_err {
		return false
	}

	_, err := os.Stat(sitepath)

	if os.IsNotExist(err) {
		return false
	}

	return true
}
