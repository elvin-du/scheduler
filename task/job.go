package task

import (
	"fmt"
	"net/http"
	"time"

	//	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

//固定时间任务
type OneTimeJob struct {
	//	Id          string
	Retry   int //失败后重试次数，默认不重试
	Target  *Target
	StartAt time.Time
	stop    chan error
	retErr  error
}

/*
失败并重试Retry次数后会把错误通过传回，
如果成功返回nil
*/

func (j *OneTimeJob) RunFunc(f func(error)) {
	j.Run()
	f(j.retErr)
}

func (j *OneTimeJob) Run() {
	to := Time2Duration(j.StartAt)

	select {
	case <-time.After(to):
		for i := 0; i < j.Retry+1; i++ {
			j.retErr = j.Do()
			if nil == j.retErr {
				break
			}
		}
	case j.retErr = <-j.stop:
	}
}

func (j *OneTimeJob) Do() error {
	url := FullURL(j.Target.Addr, j.Target.Path, j.Target.Params)
	resp, err := http.Get(url)
	if nil != err {
		return err
	}

	err = resp.Body.Close()
	if nil != err {
		log.Errorln(err)
		return err
	}

	if http.StatusOK != resp.StatusCode {
		return fmt.Errorf("url:%s,expected statusCode %d,but got %d", url, http.StatusOK, resp.StatusCode)
	}

	return nil
}

/*
只是关闭本次任务的进程，但没有彻底关闭任务，下次重启本服务，还会从数据库中取出此任务，开始继续执行
*/
func (j *OneTimeJob) Stop(reason string) {
	var err error = nil
	if "" != reason {
		err = fmt.Errorf("%s", reason)
	}

	j.stop <- err
}

func (j *OneTimeJob) Restart() {
	j.Stop("restart")
	j.reload()
	j.Run()
}

func (j *OneTimeJob) RestartFunc(f func(error)) {
	j.Stop("restart")
	j.reload()
	j.RunFunc(f)
}

func (j *OneTimeJob) reload() {

}

//循环定时任务
type EveryJob struct {
	Retry  int //失败后重试次数，默认不重试
	Target *Target
	Spec   string
}
