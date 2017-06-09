package task

import (
	"scheduler/dal"
	"time"

	log "github.com/sirupsen/logrus"
)

func StartAll() {
	jobs, err := dal.OneTimeJobModel().GetAll()
	if nil != err {
		log.Errorln(err)
		return
	}

	jobs2 := CreateJobs(jobs)
	for _, j := range jobs2 {
		go j.Run()
	}
}

func CreateJobs(js []*dal.OneTimeJob) []*OneTimeJob {
	jobs := []*OneTimeJob{}
	for _, v := range js {
		jobs = append(jobs, CreateJob(v))
	}

	return jobs
}

func CreateJob(j *dal.OneTimeJob) *OneTimeJob {
	return &OneTimeJob{
		Retry: j.Retry,
		Target: &Target{
			Addr:   j.Addr,
			Path:   j.Path,
			Params: StringToMap(j.Params),
		},

		StartAt:   time.Unix(j.StartAt, 0),
		stop:      make(chan error),
		retErr:    nil,
	}
}
