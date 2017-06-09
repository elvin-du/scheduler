package dal

import (
	"gokit/config"
	"gokit/db"

	log "github.com/sirupsen/logrus"
)

func init() {
	var (
		addr, user, passwd, dbname string
		pool                       int
	)

	err := config.Get("mysql:main:addr", &addr)
	if nil != err {
		log.Fatalln(err)
	}

	err = config.Get("mysql:main:user", &user)
	if nil != err {
		log.Fatalln(err)
	}

	err = config.Get("mysql:main:passwd", &passwd)
	if nil != err {
		log.Fatalln(err)
	}

	err = config.Get("mysql:main:dbname", &dbname)
	if nil != err {
		log.Fatalln(err)
	}

	err = config.Get("mysql:main:pool", &pool)
	if nil != err {
		log.Fatalln(err)
	}

	db.Init(addr, user, passwd, dbname, pool)
}
