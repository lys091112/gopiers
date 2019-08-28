package database

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type AlertEvent struct {
	UnqId         string
	CycleId       int64
	Key           string
	RuleId        string
	EventCategory string
	Metrics       map[string]string
	AlertSeverity string
	Timestamp     int64
	TimeScope     int64
	Time          int64
	Tags          map[string]string
}

func Generator(ruleId, key, category, severity string, metrics, tags map[string]string, cycleId int64) *AlertEvent {
	time := time.Now().Unix()
	event := &AlertEvent{
		UnqId:         uuid.New().String(),
		CycleId:       cycleId,
		Key:           key,
		RuleId:        ruleId,
		EventCategory: category,
		Metrics:       metrics,
		AlertSeverity: severity,
		Timestamp:     time,
		TimeScope:     time / 1000,
		Time:          time,
		Tags:          tags}
	return event
}

func (event AlertEvent) Insert() bool {
	db, error := GetConn()
	defer db.Close()
	CheckErr(error)
	sql := "insert into events (`unq_id`,`cycle_id`,`key`,`rule_id`,`event_category`,`severity`,`time_scope`,`time`,`metrics`,`tags`,`timestamp`) " +
		"values (?,?,?,?,?,?,?,?,?,?,?)"

	stat, err := db.Prepare(sql)
	defer stat.Close()
	CheckErr(err)

	startTime := time.Now().Unix()
	fmt.Printf("start time: %d\n", startTime)

	num, err := process(stat,&event)
	CheckErr(err)

	fmt.Printf("persist time: %d", time.Now().Unix()-startTime)
	return num > 0
}

func DoInsert() bool {
	db, error := GetConn()

	defer db.Close()
	CheckErr(error)
	sql := "insert into events (`unq_id`,`cycle_id`,`key`,`rule_id`,`event_category`,`severity`,`time_scope`,`time`,`metrics`,`tags`,`timestamp`) " +
		"values (?,?,?,?,?,?,?,?,?,?,?)"

	stat, err := db.Prepare(sql)
	defer stat.Close()
	CheckErr(err)

	startTime := time.Now().Unix()
	fmt.Printf("start time: %d\n", startTime)

	r := rand.New(rand.NewSource(time.Now().Unix()))
	tmp := make(map[string]string)
	for i := 0; i < 200; i++ {
		createTime := time.Now().Unix()
		key := uuid.New().String()
		for j := 0; j < 10000; j++ {
			ruleId := "ALERT_AI_" + strconv.FormatInt(createTime, 10) + "_nodeType_Agent_" + strconv.Itoa(r.Intn(100))
			event := Generator(ruleId, key, "HealthStatus", "Warning", tmp, tmp, int64(r.Intn(100)))
			_,err = process(stat,event)
			CheckErr(err)
		}
		fmt.Println("------------------------> " + strconv.Itoa(i))
	}
	fmt.Printf("end time: %d \n", time.Now().Unix())
	fmt.Printf("persist time: %d", time.Now().Unix()-startTime)
	return true
}

func process(stat *sql.Stmt, event *AlertEvent) (int64, error) {
	res, err := stat.Exec(event.UnqId, event.CycleId, event.Key, event.RuleId, event.EventCategory,
		event.AlertSeverity, event.TimeScope, event.Time, "{}", "{}", event.Timestamp)
	CheckErr(err)
	return res.RowsAffected()
}
