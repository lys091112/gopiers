package database

import (
	"time"

	"github.com/google/uuid"
)

type alertEvent struct {
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

func Generator(ruleId, category, severity string, metrics, tags map[string]string) *alertEvent {
	time := time.Now().Unix()
	event := &alertEvent{
		Key:           uuid.New().String(),
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

func New(time int64, ruleId string) *alertEvent {
	return &alertEvent{
		Key:    uuid.New().String(),
		Time:   time,
		RuleId: ruleId}
}
