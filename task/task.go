package task

import (
	"fmt"
	"net/http"
	"time"

	//	"github.com/robfig/cron"
)

//固定时间任务
type FixedJob struct {
	Id        string
	Name      string
	Retry     int //失败后重试次数，默认不重试
	Target    *Target
	StartAt   time.Time
	CreatedAt time.Time
	stop      chan error
	retErr    error
}

/*
失败并重试Retry次数后会把错误通过传回，
如果成功返回nil
*/

func (fj *FixedJob) RunFunc(f func(error)) {
	fj.Run()
	f(fj.retErr)
}

func (fj *FixedJob) Run() {
	to := Time2Duration(fj.StartAt)

	select {
	case <-time.After(to):
		for i := 0; i < fj.Retry+1; i++ {
			fj.retErr = fj.Do()
			if nil == fj.retErr {
				break
			}
		}
	case fj.retErr = <-fj.stop:
	}
}

func (fj *FixedJob) Do() error {
	url := FullURL(fj.Target.Addr, fj.Target.Path, fj.Target.Params)
	resp, err := http.Get(url)
	if nil != err {
		return err
	}

	err = resp.Body.Close()
	if nil != err {
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
func (fj *FixedJob) Stop(reason string) {
	var err error = nil
	if "" != reason {
		err = fmt.Errorf("%s", reason)
	}

	fj.stop <- err
}

//循环定时任务
type EveryJob struct {
	Id        string
	Name      string
	Retry     int //失败后重试次数，默认不重试
	Target    *Target
	Spec      string
	CreatedAT time.Time
}
