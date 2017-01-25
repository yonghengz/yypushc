package yypush

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/samuel/go-zookeeper/zk"
)

type Context struct {
	Zk        *zk.Conn
	Path      string
	Action    string
	Args      string
	Configure *Configure
	Stat      *zk.Stat
}

func (c *Context) IpAddDel() error {
	if c.Args == "" {
		return fmt.Errorf("action %s requires IP list", c.Action)
	}
	_, err := GetConfiguration(c)
	if err != nil {
		return err
	}
	var ips = map[string]bool{}
	ps := strings.Split(c.Configure.Ips, ",")
	nps := strings.Split(c.Args, ",")
	for _, ip := range ps {
		if ip != "" {
			ips[ip] = true
		}
	}
	if c.Action == "ipadd" {
		for _, ip := range nps {
			if ip != "" {
				ips[ip] = true
			}
		}
	} else {
		for _, ip := range nps {
			if _, ok := ips[ip]; ok {
				delete(ips, ip)
			}
		}
	}
	var allIps = []string{}
	for ip, _ := range ips {
		allIps = append(allIps, ip)
	}
	log.Printf("Total IPs: %d", len(allIps))
	c.Configure.Ips = strings.Join(allIps, ",")
	err = c.Save()
	if err != nil {
		return err
	}
	s, err := encodeConfI(c.Configure)
	if err != nil {
		return err
	}
	fmt.Printf("%s", string(s))
	return nil
}

func (c *Context) GetConf() error {
	conf, err := GetConfiguration(c)
	if err != nil {
		return err
	}
	s, err := encodeConfI(conf)
	if err != nil {
		return err
	}
	fmt.Printf("%s", string(s))
	return nil
}

func encodeConfI(conf *Configure) ([]byte, error) {
	s, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		return nil, err
	}
	return s, nil
}

func encodeConf(conf *Configure) ([]byte, error) {
	s, err := json.Marshal(conf)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *Context) Save() error {
	j, err := encodeConf(c.Configure)
	if err != nil {
		return err
	}
	//log.Printf("Before Save: Old Version: %d", c.Stat.Version)
	//s, err := c.Zk.Set(c.Path, j, c.Stat.Version+1)
	s, err := c.Zk.Set(c.Path, j, -1)
	if err != nil {
		return err
	}
	log.Printf("Saved: Old Version: %d, New Version: %d", c.Stat.Version, s.Version)
	c.Stat = s
	return nil
}
