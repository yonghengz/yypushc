package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/samuel/go-zookeeper/zk"
	"github.com/uveio/yypushc/yypush"
)

func main() {
	zkc := os.Getenv("ZOOKEEPER")
	if zkc == "" {
		log.Printf("Error: environment variable ZOOKEEPER not found")
		os.Exit(1)
	}
	log.Printf("Zookeepers: %s", zkc)
	//zks := []string{"10.13.2.43:2181", "10.13.40.53:2181"}
	zks := strings.Split(zkc, ",")
	if len(zks) == 0 {
		log.Printf("Error: failed to retrive zookeeper servers")
		os.Exit(5)
	}

	var ctx = &yypush.Context{}
	err := parseArgs(ctx)

	if err != nil {
		log.Printf("Error: %s", err)
		os.Exit(10)
	}

	c, _, err := zk.Connect(zks, time.Second) //*10)
	if err != nil {
		log.Printf("Error: %s", err)
		os.Exit(15)
	}
	ctx.Zk = c
	err = process(ctx)
	if err != nil {
		log.Printf("Error: %s", err)
		os.Exit(20)
	}
}

func parseArgs(ctx *yypush.Context) error {
	if len(os.Args) < 3 {
		return fmt.Errorf("action and path are required")
	}
	ctx.Action = os.Args[1]
	ctx.Path = os.Args[2]
	if len(os.Args) > 3 {
		ctx.Args = os.Args[3]
	}
	return nil
}

func process(ctx *yypush.Context) error {
	switch ctx.Action {
	case "ipadd", "ipdel":
		return ctx.IpAddDel()
	case "get":
		return ctx.GetConf()
	default:
		return fmt.Errorf("Not implemented action: %s", ctx.Action)
	}
	return nil
}
