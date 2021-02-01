package files

import (
	"bytes"
	"encoding/csv"
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