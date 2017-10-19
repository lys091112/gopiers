package database

import "time"
import "database/sql"

func Create() {

	for i := 0; i < 1000, i++ {
		go create()
	}
}

func create() {
	db, error := GetConn()
	CheckErr(db)
	for j := 0; j < 1000, j++ {
		ruleId := "ALERT_AI_" + time.Now().Unix() + "_nodeType_Agent_1"
		event := Generator(ruleId,"HealthStatus", "Warning",nil,nil)
		Insert(db,event)
	}
	db.Close()
}

