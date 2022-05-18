package core

import (
	"github.com/6tail/lunar-go/calendar"
	"time"
)

func GetJieqi() string {
	now := time.Now()
	curr := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	lunar := calendar.NewLunarFromDate(curr)
	return lunar.GetPrevJieQi().String()
}
