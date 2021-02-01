package card

import (
	"errors"
	"fmt"
	"github.com/Geniuskaa/Task7.1_BGO-3/pkg/transaction"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type Card struct {
	Id int64
	Issuer string
	Currency string
	Balance int64
	Number string
	Transactions []*transaction.Transaction
}

type Service struct {
	bank string
	StoreOfCards []*Card
}

type part struct {
	monthTimestamp int64
	transactions []*transaction.Transaction
}

func NewService(storeOfCards []*Card, bankName string) *Service {
	return &Service{
		bank: bankName,
		StoreOfCards: storeOfCards}
}

func (s *Service) AddCard(id int64, issuer string, currency string, balance int64, number string) {
	s.StoreOfCards = append(s.StoreOfCards, &Card{
		Id:       id,
		Issuer:   issuer,
		Currency: currency,
		Balance:  balance,
		Number:   number,
	})
}

var ErrCardNotInOurBase = errors.New("Данной карты нет в нашей базе данных.")


func (s *Service) SearchCards(number string) (err error, index int) {
	for i, _ := range s.StoreOfCards {
		if s.StoreOfCards[i].Number == number {
			return nil, i
		}
	}
	if strings.HasPrefix(number, "5106 21") {
		s.StoreOfCards = append(s.StoreOfCards, &Card{
			Id:           rand.Int63n(1000),
			Issuer:       "VISA",
			Currency:     "RUB",
			Balance:      rand.Int63n(10000000),
			Number:       number,
			Transactions: nil,
		})
		return nil, len(s.StoreOfCards) - 1
	}
	return ErrCardNotInOurBase, -1
}

func sum(transactions []*transaction.Transaction) int64 {
	result := int64(0)
	for _, transaction := range transactions {
		result += transaction.Amount
	}
	return result
}

func (t *Card) SumConcurrently(goroutines int, from time.Time, to time.Time) int64 {
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	months := make([]*part, 0)

	next := from
	for next.Before(to) {
		months = append(months, &part{
			monthTimestamp: next.Unix(),
		})
		next = next.AddDate(0, 1, 0)
	}
	months = append(months, &part{
		monthTimestamp: to.Unix(),
	})

	for j, transaction := range t.Transactions {
		if months[0].monthTimestamp <= transaction.Date.Unix() && transaction.Date.Unix() < months[len(months) - 1].monthTimestamp {
			for i := 1; i < len(months); i++ {
				if t.Transactions[j].Date.Unix() < months[i].monthTimestamp {
					months[i - 1].transactions = append(months[i - 1].transactions, t.Transactions[j])
					break
				}
			}
		}
	}

	total := int64(0)
	partSize := len(months) / goroutines // Динамически поменяем в For
	for i := 0; i < goroutines; i++ {

		part := months[i*partSize : (i+1)*partSize]
		go func() {
			for _, element := range part {
				sum := sum(element.transactions)
				total += sum
				fmt.Println(sum / 100)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("За выбранный промежуток времени было потрачено ", total / 100, " рублей")
	return total
}

//func mccChecker(MCC string) string {
//	switch MCC {
//	case "5090":
//		return "Финансы"
//	case "5050":
//		return "Рестораны"
//	case "5105":
//		return "Транспорт"
//	case "5060":
//		return "Супермаркеты"
//	default:
//		return "Остальное"
//	}
//}


func MonthlySpendings(transactions []*transaction.Transaction) map[string]int64 {

	mcc := map[string]int64 {
		"5050": 0,
		"5090": 0,
		"5105": 0,
		"5060": 0,
	}

	for index := range transactions {
		mcc[transactions[index].MCC] += transactions[index].Amount
	}
	return mcc
}

func (c *Card) MonthlySpendingsMutex(goroutines int) map[string]int64 {
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	mu := sync.Mutex{}
	mcc := map[string]int64 {
		"5050": 0,
		"5090": 0,
		"5105": 0,
		"5060": 0,
	}
	partSize := len(c.Transactions) / goroutines
	for i := 0; i < goroutines; i++ {
		part := c.Transactions[i*partSize : (i+1)*partSize]
		go func() {
			pieceMCC := MonthlySpendings(part)
			mu.Lock()
			for key, value := range pieceMCC {
				mcc[key] += value
			}
			mu.Unlock()
			wg.Done()
		}()
	}
	lastPiece := MonthlySpendings(c.Transactions[goroutines*partSize : ])
	wg.Wait()
	for key, value := range lastPiece {
		mcc[key] += value
	}
	return mcc
}

func (c *Card)MonthlySpendingsChanels(goroutines int) map[string]int64 {

	mcc := map[string]int64 {
		"5050": 0,
		"5090": 0,
		"5105": 0,
		"5060": 0,
	}
	result := make(chan map[string]int64)
	partSize := len(c.Transactions) / goroutines
	for i := 0; i < goroutines; i++ {
		part := c.Transactions[i*partSize : (i+1)*partSize]
		monthlySpendings(part, result)
	}
	finished := 0
	for value := range result {
		for key, value2 := range value {
			mcc[key] += value2
		}
		finished++
		if finished == goroutines {
			close(result)
			break
		}
	}
	lastPiece := MonthlySpendings(c.Transactions[goroutines*partSize : ])
	for key, value := range lastPiece {
		mcc[key] += value
	}
	return mcc
}

func monthlySpendings(transactions []*transaction.Transaction, result chan <- map[string]int64) {
	go func() {
		result <- MonthlySpendings(transactions)
	}()
}

func (c *Card)MonthlySpendingsMutex2(goroutines int) map[string]int64 { // без 1-ой функции
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	mu := sync.Mutex{}
	mcc := map[string]int64 {
		"5050": 0,
		"5090": 0,
		"5105": 0,
		"5060": 0,
	}
	partSize := len(c.Transactions) / goroutines
	for i := 0; i < goroutines; i++ {
		part := c.Transactions[i*partSize : (i+1)*partSize]
		go func() {
			mu.Lock()
			for _, value := range part {
				mcc[value.MCC] += value.Amount
			}
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	for _, value := range c.Transactions[goroutines*partSize : ] {
		mcc[value.MCC] += value.Amount
	}
	return mcc
}

func PrintMapOfMCC(m map[string]int64) {
	sliceOfMcc := []int64{}
	sliceOfMcc = append(sliceOfMcc, m["5090"])
	fmt.Println("Финансы ", sliceOfMcc[0]/100, " рублей")
	sliceOfMcc = append(sliceOfMcc, m["5050"])
	fmt.Println("Рестораны ", sliceOfMcc[1]/100, " рублей")
	sliceOfMcc = append(sliceOfMcc, m["5105"])
	fmt.Println("Транспорт ", sliceOfMcc[2]/100, " рублей")
	sliceOfMcc = append(sliceOfMcc, m["5060"])
	fmt.Println("Супермаркеты ", sliceOfMcc[3]/100, " рублей")

}



