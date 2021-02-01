package files

import (
	"encoding/csv"
	"encoding/json"
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

func ExportJson(data []byte) {
	var decoded []transaction.Transaction

	err := json.Unmarshal(data, &decoded)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%#v", decoded)
}

func ExportXml(data []byte) {
	var decoded []transaction.Transactions

	err := xml.Unmarshal(data, &decoded)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%#v", decoded)
}