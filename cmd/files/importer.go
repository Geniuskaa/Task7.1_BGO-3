package files

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"github.com/Geniuskaa/Task7.1_BGO-3/pkg/card"
	"github.com/Geniuskaa/Task7.1_BGO-3/pkg/transaction"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

func Import(fileName string, card *card.Card)  {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(err)
		return
	}

	reader := csv.NewReader(bytes.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return
	}



	for _, string := range records {
		card.Transactions = append(card.Transactions, mapRowToTransaction(string))
	}
}

func mapRowToTransaction(slice []string) (*transaction.Transaction) {
	id, _ := strconv.Atoi(slice[3])
	amount, _ := strconv.Atoi(slice[2])
	date, _ := strconv.Atoi(slice[0])

	return &transaction.Transaction{
		Id:     int64(id),
		Amount: int64(amount),
		MCC:    slice[1],
		Date:   time.Unix(int64(date), 0),
		Status: "Completed",
	}
}

func ImportJson(transactions []*transaction.Transaction) []byte {
	encoded, err := json.Marshal(transactions)
	if err != nil {
		log.Println(err)
		return []byte{0}
	}

	return encoded
}

func ImportXml(transactions []*transaction.Transaction) []byte {
	var sliceOfTransactions []transaction.Transaction
	for _, element := range transactions {
		sliceOfTransactions = append(sliceOfTransactions, transaction.Transaction{
			Id:      element.Id,
			Amount:  element.Amount,
			MCC:     element.MCC,
			Date:    element.Date,
			Status:  element.Status,
		})
	}

	changedTransactions := transaction.Transactions{
		Transactions: sliceOfTransactions,
	}

	encoded, err := xml.Marshal(changedTransactions)
	if err != nil {
		log.Println(err)
		return []byte{0}
	}
	encoded = append([]byte(xml.Header), encoded...)
	log.Println(string(encoded))
	return encoded
}