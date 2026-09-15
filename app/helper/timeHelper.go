package helper

import "time"

func StringToDate(dateStr string)(time.Time,error){
	layout:="2006-01-02"
	return time.Parse(layout,dateStr)
}