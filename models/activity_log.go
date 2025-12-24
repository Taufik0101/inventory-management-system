package models

import "time"

type ActivityLog struct {
	ID        int       `json:"id"`
	ItemID    int       `json:"item_id"`
	Action    string    `json:"action"`
	Amount    int       `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

var Logs = []ActivityLog{}
var NextLogID = 1
