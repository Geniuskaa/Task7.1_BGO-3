package files

import (
	"encoding/csv"
	"encoding/xml"
	"github.com/Geniuskaa/Task7.1_BGO-3/pkg/transaction"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

type Service struct {
	mu sync.Mutex
	transactions []*transaction.Transaction
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Export(writer io.Writer) error {
	s.mu.Lock()
	if len(s.transactions) == 0 {
		s.mu.Unlock()
		return nil
	}

	records := make([][]string, len(s.transactions))
	for _, t := range s.transactions {
		record := []string{
			strconv.Itoa(int(t.Date.Unix())),
			t.MCC,
			strconv.Itoa(int(t.Amount)),
			strconv.Itoa(int(t.Id)),
		}
		records = append(records, record)
	}
	s.mu.Unlock()

	w := csv.NewWriter(writer)
	return w.WriteAll(records)
}

func ExportTransactions(nameOfFile string, sliceOfTransactions []*transaction.Transaction) error {
	file, err := os.Create(nameOfFile)
	if err != nil {
		log.Println(err)
		return err
	}

	defer func(c io.Closer) {
		if err := c.Close(); err != nil {
			log.Println(err)
		}
	}(file)

	svc := NewService()

	for _, element := range sliceOfTransactions {
		svc.transactions = append(svc.transactions, element)
	}

	err = svc.Export(file)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func ExportXml(nameOfFile string, transactions []*transaction.Transaction) error {
	name := nameOfFile + ".xml"
	file, err := os.Create(name)
	if err != nil {
		log.Println(err)
		return err
	}

	defer func(c io.Closer) {
		if err := c.Close(); err != nil {
			log.Println(err)
		}
	}(file)

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
		return err
	}
	encoded = append([]byte(xml.Header), encoded...)


	w := xml.NewEncoder(file)
	w.Encode(string(encoded))
	return nil
}