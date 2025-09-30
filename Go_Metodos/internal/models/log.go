package models

import "time"

//informações necessárias para auditoria de items
type Log struct {
	TimeStamp time.Time
	Action    string
	User      string
	ItemId    int
	Quantity  int
	Reason    string
}
