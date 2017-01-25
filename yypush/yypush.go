package yypush

import (
	"encoding/json"

	"github.com/samuel/go-zookeeper/zk"
)

type Configure struct {
	BasePath  string `json:"basePath"`
	Suffix    string `json:"suffix"`
	KafkaName string `json:"kafkaName"`
	Ips       string `json:"ips"`
	//StartTime          time.Time  `json:"startTime"`
	StartTime          string     `json:"startTime"`
	Feed               string     `json:"feed"`
	Delimiter          string     `json:"delimiter"`
	FileType           string     `json:"fileType"`
	UseFileCurrent     string     `json:"useFileCurrent"`
	KafkaProducerProps KafkaProps `json:"kafkaProducerProps"`
}

type KafkaProps struct {
	BootstrapServers string `json:"bootstrap.servers"`
	Acks             string `json:"acks"`
	Retries          int    `json:"retries"`
	BatchSize        int    `json:"batch.size"`
	LingerMs         int    `json:"linger.ms"`
	BufferMemory     int    `json:"buffer.memory"`
	KeySerializer    string `json:"key.serializer"`
	ValueSerializer  string `json:"value.serializer"`
}

func GetConfiguration(ctx *Context) (*Configure, error) {
	v, s, err := zkGet(ctx)
	if err != nil {
		return nil, err
	}
	var conf Configure
	err = json.Unmarshal(v, &conf)
	if err != nil {
		return nil, err
	}
	ctx.Configure = &conf
	ctx.Stat = s
	return &conf, nil
}

func SaveConfiguration(ctx *Context) error {
	return nil
}

func zkGet(ctx *Context) ([]byte, *zk.Stat, error) {
	v, s, err := ctx.Zk.Get(ctx.Path)
	if err != nil {
		return nil, nil, err
	}
	return v, s, nil
}
