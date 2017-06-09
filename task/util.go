package task

import (
	"strings"
	"time"
	//	"encoding/json"
	json "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

func Time2Duration(t time.Time) time.Duration {
	return t.Sub(time.Now())
}

func StringToDuration(s string) (time.Duration, error) {
	t, err := time.Parse("yyyy-MM-dd HH:mm:ss", s)
	if nil != err {
		return 0, err
	}

	return Time2Duration(t), nil
}

func FullURL(addr, path string, params map[string]string) string {
	url := addr + path + "?"
	for k, v := range params {
		url += k + "=" + v + "&"
	}
	return strings.TrimSuffix(url, "&")
}

func StringToMap(s string) map[string]string {
	m := make(map[string]string)
	err := json.UnmarshalFromString(s, &m)
	if nil != err {
		log.Errorln(err)
	}

	return m
}
