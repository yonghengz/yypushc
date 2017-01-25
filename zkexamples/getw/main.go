package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	zks := []string{"10.13.2.43:2181", "10.13.40.53:2181"}
	c, _, err := zk.Connect(zks, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	slog, stat, ch, err := c.GetW("/logpush/stat")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v %+v\n", string(slog), stat)
	e := <-ch
	fmt.Printf("%+v\n", e)
}
