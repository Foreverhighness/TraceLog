package es

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/kataras/golog"
	"github.com/olivere/elastic"
)

var (
	client   *elastic.Client
	dataChan chan *Parameter
)

// Parameter ...
type Parameter struct {
	Index string
	Data  string
}

// Init ...
func Init(addrs string, chansize int) (err error) {
	if nil != client {
		return nil
	}
	if !strings.HasPrefix(addrs, "http://") {
		addrs = "http://" + addrs
	}
	client, err = elastic.NewClient(elastic.SetURL(addrs))
	if err != nil {
		golog.Error("[es] Fail to connect ES. err: ", err)
		return
	}
	info, code, err := client.Ping(addrs).Do(context.Background())
	if err != nil {
		golog.Error("[es] ES do not response. err: ", err)
		return
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	dataChan = make(chan *Parameter, chansize)
	go sendToES()
	return
}

// SendToChan ...
func SendToChan(data *Parameter) {
	dataChan <- data
}

func sendToES() {
	for {
		select {
		case message := <-dataChan:
			put, err := client.Index().
				Index(message.Index).
				BodyString(message.Data).
				Do(context.Background())
			if err != nil {
				golog.Error("[es] Fail to put message to es. err: ", err)
				return
			}
			fmt.Printf("Indexed tweet %s to index %s, type %s\n", put.Id, put.Index, put.Type)
		default:
			time.Sleep(50 * time.Millisecond)
		}

	}
}
