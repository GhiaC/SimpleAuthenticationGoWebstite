package Controler

import (
	"github.com/go-xorm/xorm"
	"fmt"
	"Website/Models"
	"log/syslog"
	"log"
)

var engine *xorm.Engine
var flag bool
var u = &Models.User{}

const userDB = "root"
const passDB  = "mghiasi"
const nameDB  = "authGo"

func GetEngine() *xorm.Engine{
	if !flag {
		flag = true
		var errDB error
		engine, errDB = xorm.NewEngine("mysql", userDB+":"+passDB+"@/"+nameDB+"?charset=utf8")
		if errDB != nil {
			fmt.Println(errDB)
		}
		engine.CreateTables(u)
		engine.Sync2(new(Models.User))
		logWriter, err := syslog.New(syslog.LOG_DEBUG, "rest-xorm-example")
		if err != nil {
			log.Fatalf("Fail to create xorm system logger: %v\n", err)
		}

		logger := xorm.NewSimpleLogger(logWriter)
		logger.ShowSQL(true)
		engine.SetLogger(logger)
	}
	return engine
}
