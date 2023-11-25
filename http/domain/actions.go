package domain

import (
	"L2/http/models"
	time2 "time"
)

type ev models.Event

func NewEvent(date, time string, userId int) *models.Event {
	Uid := time2.Now().Unix()
	return &models.Event{
		UserId: userId,
		Date:   date,
		Time:   time,
		Uid:    Uid,
	}
}

func (e *ev) UpdateEvent(newDate, newTime string) {
	if len(newDate) != 0 {
		e.Date = newDate
	}
	if len(newTime) != 0 {
		e.Time = newTime
	}
}
