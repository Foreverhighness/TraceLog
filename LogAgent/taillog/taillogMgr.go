package taillog

import (
	"errors"
	"fmt"
	"logagent/kafka"
	"time"

	"github.com/kataras/golog"
)

// TailTaskConfig to control tailtask
type TailTaskConfig struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

type taillogMgr struct {
	taskMap       map[string]*tailTask
	configs       []*TailTaskConfig
	updateChannel chan []*TailTaskConfig
}

func closure(topic string) func(string) {
	return func(line string) {
		kafka.SendToChan(topic, line)
	}
}

var instance *taillogMgr

// Init by testEntry
func Init(entrys []*TailTaskConfig) (err error) {
	if instance != nil {
		return errors.New("Duplicate init")
	}
	instance = &taillogMgr{
		taskMap:       make(map[string]*tailTask, 16),
		configs:       entrys,
		updateChannel: make(chan []*TailTaskConfig),
	}
	for _, entry := range entrys {
		err = addTask(entry)
		if err != nil {
			golog.Error("[tailMgr] Fail to add task. err:", err)
			return
		}
	}
	go listenUpdate()
	return
}

func addTask(entry *TailTaskConfig) (err error) {
	var task *tailTask
	taskinfo := fmt.Sprintf("%s_%s", entry.Path, entry.Topic)
	handle := closure(entry.Topic)
	task, err = newTask(entry.Path, taskinfo, handle)
	if err != nil {
		golog.Error("[tailMgr] Fail to generate task. err:", err)
		return
	}
	instance.taskMap[taskinfo] = task
	go task.run()
	return
}

func listenUpdate() {
	var err error
	for {
		select {
		case newEntrys := <-instance.updateChannel:
			oldEntrys := instance.configs
		addLoop:
			for _, newEntry := range newEntrys {
				newTaskinfo := fmt.Sprintf("%s_%s", newEntry.Path, newEntry.Topic)
				for _, oldEntry := range oldEntrys {
					oldTaskinfo := fmt.Sprintf("%s_%s", oldEntry.Path, oldEntry.Topic)
					if newTaskinfo == oldTaskinfo {
						continue addLoop
					}
				}
				err = addTask(newEntry)
				if err != nil {
					golog.Error("[tailMgr] Fail to update task. err:", err)
					return
				}
			}
		deleteLoop:
			for _, oldEntry := range oldEntrys {
				oldTaskinfo := fmt.Sprintf("%s_%s", oldEntry.Path, oldEntry.Topic)
				for _, newEntry := range newEntrys {
					newTaskinfo := fmt.Sprintf("%s_%s", newEntry.Path, newEntry.Topic)
					if oldTaskinfo == newTaskinfo {
						continue deleteLoop
					}
				}
				instance.taskMap[oldTaskinfo].cancel()
			}
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// GetUpdateChan Get channel
func GetUpdateChan() chan<- []*TailTaskConfig {
	return instance.updateChannel
}
