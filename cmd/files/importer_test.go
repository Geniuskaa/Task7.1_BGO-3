package files

import (
	"github.com/Geniuskaa/Task7.1_BGO-3/pkg/transaction"
	"reflect"
	"testing"
	"time"
)

func Test_mapRowToTransaction(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name string
		args args
		want *transaction.Transaction
	}{
		{"Fisrt", args{slice: []string{"1614135600","5090","1314600","20"}}, &transaction.Transaction{
			Id:     20,
			Amount: 1314600,
			MCC:    "5090",
			Date:   time.Date(2021, 2,24,6,0,0,0,time.Local),
			Status: "Completed",
		}},

		{"Second", args{slice: []string{"1618369200","5050","7341600","20"}}, &transaction.Transaction{
			Id:     20,
			Amount: 7341600,
			MCC:    "5050",
			Date:   time.Date(2021, 4,14,6,0,0,0,time.Local),
			Status: "Completed",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapRowToTransaction(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapRowToTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}