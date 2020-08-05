package taillog

import (
	"context"
	"time"

	"github.com/hpcloud/tail"
	"github.com/kataras/golog"
)

type statusType uint32

const (
	initialStatus statusType = iota
	runningStatus
	pauseStatus
	haltStatus
	invaildStatus
)

type tailTask struct {
	Status   statusType
	TaskInfo string
	tailObj  *tail.Tail
	handle   func(string)
	ctx      context.Context
	cancel   context.CancelFunc
}

// newTask something

func newTask(filename, taskInfo string, handle func(string)) (task *tailTask, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	task = &tailTask{
		Status:   initialStatus,
		TaskInfo: taskInfo,
		handle:   handle,
		ctx:      ctx,
		cancel:   cancel,
	}
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件那个位置开始读
		MustExist: false,                                // 文件不存在报错
		Poll:      true,                                 // 是否轮询
	}
	task.tailObj, err = tail.TailFile(filename, config)
	if err != nil {
		golog.Error("[tailTask] Tail file failed, err:%v\n", err)
		return
	}
	task.Status = pauseStatus
	return
}

func (tsk *tailTask) run() {
	for {
		select {
		case <-tsk.ctx.Done():
			tsk.Status = haltStatus
			golog.Infof("[tailTask] <%s> is over.", tsk.TaskInfo)
			return
		case line, ok := <-tsk.tailObj.Lines:
			if !ok {
				golog.Error("[tailTask] Tail file close reopen, filename:%s\n", tsk.tailObj.Filename)
				time.Sleep(time.Second)
				continue
			}
			golog.Debug("[tailTask] Get line: ", line.Text)
			tsk.handle(line.Text)
		default:
			time.Sleep(50 * time.Millisecond)
		}

	}
}
