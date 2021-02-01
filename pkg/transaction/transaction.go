package transaction

import "time"

type Transaction struct {
	XMLName string `xml:"transaction"`
	Id int64 `xml:"id"`
	Amount int64 `xml:"amount"`
	MCC string `xml:"mcc"`
	Date time.Time `xml:"date"`
	Status string `xml:"status"`
}

type Transactions struct {
	XMLName string `xml:"transactions"`
	Transactions []Transaction
}






