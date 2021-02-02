package main

import (
	"fmt"
	"github.com/Geniuskaa/Task7.1_BGO-3/cmd/files"
	"github.com/Geniuskaa/Task7.1_BGO-3/pkg/card"
	"github.com/Geniuskaa/Task7.1_BGO-3/pkg/transaction"
	"github.com/Geniuskaa/Task7.1_BGO-3/pkg/transfer"
	"io"
	"log"
	"os"
	"runtime/trace"
	"sort"
	"time"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Print(err)
		}
	}()
	err = trace.Start(f)
	if err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()



	bank := card.NewService([]*card.Card{},"Tinkoff")
	bank.AddCard(1,"VISA", "RUB", 14_800_00, "4724 3728 3929 5030")
	bank.AddCard(2, "MASTER", "RUB", 28_750_00, "6930 2857 3892 2967")
	bank.AddCard(3, "VISA", "RUB", 352_362_00, "4626 9205 2859 2852")


	transfers := transfer.NewService(bank, 0, 0.5, 10_00, 1.5, 30_00)
	_, err = transfers.Card2Card("4724 3728 3929 5030", "6930 2857 3892 2967", 5_425, "5090", time.Now())
	if err != nil {
		switch err {
		case transfer.ErrMoneyOnCardOfSenderDontEnough:
			fmt.Println("Недостаточно средств на балансе для перевода.")
		case transfer.ErrTooLowSumOfTransfer:
			fmt.Println("Слишком маленькая сумма перевода.")
		default:
			fmt.Println("Возникла непредвиденная ошибка.")
		}
	}

	transfers.Purchase(1_204, 0, "5050", time.Date(2021,2,14,6,0,0,0, time.Local))
	transfers.Purchase(13_146, 0, "5090", time.Date(2021,2,24,6,0,0,0, time.Local))
	transfers.Purchase(106, 0, "5105", time.Date(2021,2,25,6,0,0,0, time.Local))
	transfers.Purchase(746, 0, "5060", time.Date(2021,3,14,6,0,0,0, time.Local))
	transfers.Purchase(2_546, 0, "5090", time.Date(2021,3,4,6,0,0,0, time.Local))
	transfers.Purchase(73_416, 0, "5050", time.Date(2021,4,14,6,0,0,0, time.Local))
	transfers.Purchase(713_416, 0, "5090", time.Date(2021,4,14,6,0,0,0, time.Local))

	if err := execute("newText2", bank.StoreOfCards[0].Transactions); err != nil {
		os.Exit(1)
	}

	//for _, sample := range bank.StoreOfCards[0].Transactions {
	//	fmt.Println(sample)
	//}
	//fmt.Println(" ")

	//SortSumOfTransactions(bank.StoreOfCards[0].Transactions)

	//for _, sample := range bank.StoreOfCards[0].Transactions {
	//	fmt.Println(sample)
	//}



	bank.StoreOfCards[0].SumConcurrently(5, time.Date(2021,1,1,0,0,0,0, time.Local), time.Date(2021,5,1,0, 0,0,0, time.Local))

	fmt.Println("Длина слайса: ", len(bank.StoreOfCards[0].Transactions))
	fmt.Println("")
	m := card.MonthlySpendings(bank.StoreOfCards[0].Transactions)
	card.PrintMapOfMCC(m)
	fmt.Println("")

	m1 := bank.StoreOfCards[0].MonthlySpendingsMutex(6)
	card.PrintMapOfMCC(m1)
	fmt.Println("")

	m2 := bank.StoreOfCards[0].MonthlySpendingsChanels(5)
	card.PrintMapOfMCC(m2)
	fmt.Println("")

	m3 := bank.StoreOfCards[0].MonthlySpendingsMutex2(5)
	card.PrintMapOfMCC(m3)

	files.ExportTransactions("newText", bank.StoreOfCards[0].Transactions)

	files.Import("newText.csv", bank.StoreOfCards[0])

	fmt.Println(bank.StoreOfCards[0].Transactions[8])




}


func execute(filename string, sliceOfTransactions []*transaction.Transaction) (err error) {
	name := filename + ".csv"
	file, err := os.Create(name)
	if err != nil {
		log.Println(err)
		return err
	}
	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			log.Println(cerr)
			if err == nil {
				err = cerr
			}
		}
	}(file)

	err = files.ExportTransactions(filename, sliceOfTransactions)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SortSumOfTransactions(transactions []*transaction.Transaction) []*transaction.Transaction {
	sort.SliceStable(transactions, func(i, j int) bool {
		return transactions[i].Amount > transactions[j].Amount
	})
	return transactions
}
