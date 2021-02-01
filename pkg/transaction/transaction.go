package transaction

import "time"

type Transaction struct {
	Id int64
	Amount int64
	MCC string
	Date time.Time
	Status string
}




