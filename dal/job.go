package dal

import (
	"fmt"
	"gokit/db"

	log "github.com/sirupsen/logrus"
)

var _ontTimeJobModel = &oneTimeJobModel{}

type oneTimeJobModel struct{}

func OneTimeJobModel() *oneTimeJobModel {

	return _ontTimeJobModel
}

func (j *oneTimeJobModel) Get(id string) (*OneTimeJob, error) {
	var val = &OneTimeJob{}
	sqlStr := fmt.Sprintf("SELECT * from one_time_jobs WHERE id = %d", id)
	err := db.MainDB().GetOne(val, sqlStr)
	if nil != err {
		log.Errorln(err)
		return nil, err
	}

	return val, nil
}

func (j *oneTimeJobModel) GetAll() ([]*OneTimeJob, error) {
	vals := []*OneTimeJob{}
	sqlStr := fmt.Sprintf("SELECT * from one_time_jobs")
	err := db.MainDB().GetAll(&vals, sqlStr)
	if nil != err {
		log.Errorln(err)
		return nil, err
	}

	return vals, nil
}
