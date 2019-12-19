package app

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

var (
	DBType string
	DBURL  string
	Addr   string
	Debug  bool
)

func init() {
	Provide(connDB)
}

func connDB() (*gorm.DB, error) {
	if DBURL == "" {
		return nil, fmt.Errorf("require dburl.")
	}
	if DBType == "mysql" {
		logrus.Debugf("try to connect mysql db %s", DBURL)
		db, err := gorm.Open("mysql", DBURL)
		if err != nil {
			logrus.Errorf("connect to mysql %s failed, %s", DBURL, err)
			return nil, fmt.Errorf("mysql connected failed, %s", err)
		}
		return db, nil
	}

	logrus.Debugf("try to connect postgres db %s", DBURL)
	db, err := gorm.Open("postgres", DBURL)
	if err != nil {
		logrus.Errorf("connect to postgres %s failed, %s", DBURL, err)
		return nil, fmt.Errorf("postgres connected failed, %s", err)
	}
	return db, nil
}
