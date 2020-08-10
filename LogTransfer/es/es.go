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
	Data  interface{}
}

// Init ...
func Init(addr string, chansize int) (err error) {
	if nil != client {
		return nil
	}
	if !strings.HasPrefix(addr, "http://") {
		addr = "http://" + addr
	}
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(addr))
	if err != nil {
		golog.Errorf("[es] Fail to connect ES by %s. err: %v", addr, err)
		return
	}
	info, code, err := client.Ping(addr).Do(context.Background())
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
				Index(strings.ToLower(message.Index)).
				Type("log").
				BodyJson(message.Data).
				Do(context.Background())
			if err != nil {
				golog.Error("[es] Fail to put message to es. err: ", err)
				return
			}
			golog.Debugf("Indexed message <%s> to index <%s>, type <%s>, %s\n", put.Id, put.Index, put.Type, put.Result)
		default:
			time.Sleep(50 * time.Millisecond)
		}

	}
}
