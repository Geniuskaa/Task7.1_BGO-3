package card

import (
	"github.com/Geniuskaa/Task7.1_BGO-3/pkg/transaction"
	"reflect"
	"testing"
	"time"
)

func makeTransactions() []*transaction.Transaction {
	const users = 1_000
	const transactionsPerUser = 1_000
	const transactionAmount = 2_00
	transactions := make([]*transaction.Transaction, users * transactionsPerUser)
	for index := range transactions {
		switch index % 100 {
		case 0:
			transactions[index] = &transaction.Transaction{
				Id:     0,
				Amount: transactionAmount, // 2 000 000
				MCC:    "5050", // рестораны
				Date:   time.Date(2020,1,1,0,0,0,0,time.Local),
				Status: "Completed",
			}
		case 5:
			transactions[index] = &transaction.Transaction{
				Id:     0,
				Amount: transactionAmount, // 2 000 000
				MCC:    "5090", // финансы
				Date:   time.Date(2020,1,1,0,0,0,0,time.Local),
				Status: "Completed",
			}
		case 20:
			transactions[index] = &transaction.Transaction{
				Id:     0,
				Amount: transactionAmount, // 2 000 000
				MCC:    "5060", // супермаркеты
				Date:   time.Date(2020,1,1,0,0,0,0,time.Local),
				Status: "Completed",
			}
		default:
			transactions[index] = &transaction.Transaction{
				Id:     0,
				Amount: transactionAmount, // 194 000 000
				MCC:    "5105", // транспорт
				Date:   time.Date(2020,1,1,0,0,0,0,time.Local),
				Status: "Completed",
			}
		}
	}
	return transactions
}

func TestCard_MonthlySpendingsChanels(t *testing.T) {
	type fields struct {
		Id           int64
		Issuer       string
		Currency     string
		Balance      int64
		Number       string
		Transactions []*transaction.Transaction
	}
	type args struct {
		goroutines int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]int64
	}{
		{"First", fields{
			Id:           0,
			Issuer:       "VISA",
			Currency:     "RUB",
			Balance:      134_324_00,
			Number:       "4263 2462 2361 9682",
			Transactions: makeTransactions(),
		}, args{5}, map[string]int64{
			"5050": 2_000_000,
			"5090": 2_000_000,
			"5060": 2_000_000,
			"5105": 194_000_000,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Card{
				Id:           tt.fields.Id,
				Issuer:       tt.fields.Issuer,
				Currency:     tt.fields.Currency,
				Balance:      tt.fields.Balance,
				Number:       tt.fields.Number,
				Transactions: tt.fields.Transactions,
			}
			if got := c.MonthlySpendingsChanels(tt.args.goroutines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MonthlySpendingsChanels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_MonthlySpendingsMutex(t *testing.T) {
	type fields struct {
		Id           int64
		Issuer       string
		Currency     string
		Balance      int64
		Number       string
		Transactions []*transaction.Transaction
	}
	type args struct {
		goroutines int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]int64
	}{
		{"First", fields{
			Id:           0,
			Issuer:       "VISA",
			Currency:     "RUB",
			Balance:      134_324_00,
			Number:       "4263 2462 2361 9682",
			Transactions: makeTransactions(),
		}, args{5}, map[string]int64{
			"5050": 2_000_000,
			"5090": 2_000_000,
			"5060": 2_000_000,
			"5105": 194_000_000,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Card{
				Id:           tt.fields.Id,
				Issuer:       tt.fields.Issuer,
				Currency:     tt.fields.Currency,
				Balance:      tt.fields.Balance,
				Number:       tt.fields.Number,
				Transactions: tt.fields.Transactions,
			}
			if got := c.MonthlySpendingsMutex(tt.args.goroutines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MonthlySpendingsMutex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_MonthlySpendingsMutex2(t *testing.T) {
	type fields struct {
		Id           int64
		Issuer       string
		Currency     string
		Balance      int64
		Number       string
		Transactions []*transaction.Transaction
	}
	type args struct {
		goroutines int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]int64
	}{
		{"First", fields{
			Id:           0,
			Issuer:       "VISA",
			Currency:     "RUB",
			Balance:      134_324_00,
			Number:       "4263 2462 2361 9682",
			Transactions: makeTransactions(),
		}, args{5}, map[string]int64{
			"5050": 2_000_000,
			"5090": 2_000_000,
			"5060": 2_000_000,
			"5105": 194_000_000,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Card{
				Id:           tt.fields.Id,
				Issuer:       tt.fields.Issuer,
				Currency:     tt.fields.Currency,
				Balance:      tt.fields.Balance,
				Number:       tt.fields.Number,
				Transactions: tt.fields.Transactions,
			}
			if got := c.MonthlySpendingsMutex2(tt.args.goroutines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MonthlySpendingsMutex2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonthlySpendings(t *testing.T) {
	type args struct {
		transactions []*transaction.Transaction
	}
	tests := []struct {
		name string
		args args
		want map[string]int64
	}{
		{"First", args{ makeTransactions()}, map[string]int64{
			"5050": 2_000_000,
			"5090": 2_000_000,
			"5060": 2_000_000,
			"5105": 194_000_000,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MonthlySpendings(tt.args.transactions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MonthlySpendings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMonthlySpendings(b *testing.B) {
	transactions := makeTransactions()
	want := map[string]int64{
		"5050": 2_000_000,
		"5090": 2_000_000,
		"5060": 2_000_000,
		"5105": 194_000_000,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := MonthlySpendings(transactions)
		b.StopTimer() // останавливаем таймер, чтобы время сравнения не учитывалось
		if !reflect.DeepEqual(result, want) {
			b.Fatalf("invalid result, got %v, want %v", result, want)
		}
		b.StartTimer()
	}
}

func BenchmarkMonthlySpendingsMutex(b *testing.B) {
	card := &Card{
		Id:           0,
		Issuer:       "VISA",
		Currency:     "RUB",
		Balance:      134_324_00,
		Number:       "4263 2462 2361 9682",
		Transactions: makeTransactions(),
	}
	want := map[string]int64{
		"5050": 2_000_000,
		"5090": 2_000_000,
		"5060": 2_000_000,
		"5105": 194_000_000,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := card.MonthlySpendingsMutex(5)
		b.StopTimer() // останавливаем таймер, чтобы время сравнения не учитывалось
		if !reflect.DeepEqual(result, want) {
			b.Fatalf("invalid result, got %v, want %v", result, want)
		}
		b.StartTimer()
	}
}


func BenchmarkMonthlySpendingsChanels(b *testing.B) {
	card := &Card{
		Id:           0,
		Issuer:       "VISA",
		Currency:     "RUB",
		Balance:      134_324_00,
		Number:       "4263 2462 2361 9682",
		Transactions: makeTransactions(),
	}
	want := map[string]int64{
		"5050": 2_000_000,
		"5090": 2_000_000,
		"5060": 2_000_000,
		"5105": 194_000_000,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := card.MonthlySpendingsChanels(5)
		b.StopTimer() // останавливаем таймер, чтобы время сравнения не учитывалось
		if !reflect.DeepEqual(result, want) {
			b.Fatalf("invalid result, got %v, want %v", result, want)
		}
		b.StartTimer()
	}
}


func BenchmarkMonthlySpendingsMutex2(b *testing.B) {
	card := &Card{
		Id:           0,
		Issuer:       "VISA",
		Currency:     "RUB",
		Balance:      134_324_00,
		Number:       "4263 2462 2361 9682",
		Transactions: makeTransactions(),
	}
	want := map[string]int64{
		"5050": 2_000_000,
		"5090": 2_000_000,
		"5060": 2_000_000,
		"5105": 194_000_000,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := card.MonthlySpendingsMutex2(5)
		b.StopTimer() // останавливаем таймер, чтобы время сравнения не учитывалось
		if !reflect.DeepEqual(result, want) {
			b.Fatalf("invalid result, got %v, want %v", result, want)
		}
		b.StartTimer()
	}
}


