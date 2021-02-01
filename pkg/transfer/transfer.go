package transfer

import (
	"errors"
	"fmt"
	"github.com/Geniuskaa/Task7.1_BGO-3/pkg/card"
	"github.com/Geniuskaa/Task7.1_BGO-3/pkg/transaction"
	"strconv"
	"strings"
	"time"
)

type Service struct {
	CardSvc           *card.Service
	toTinkPercent     float64
	fromTinkPercent   float64
	fromTinkMinSum    int64
	otherCardsPercent float64
	otherCardsMinSum  int64
}

func NewService(cardSvc *card.Service, toTinPer float64, fromTinPer float64, fromTinMSum int64, otherCardPer float64, otherCardMSum int64) *Service {
	return &Service{
		CardSvc:           cardSvc,
		toTinkPercent:     toTinPer,
		fromTinkPercent:   fromTinPer / 100,
		fromTinkMinSum:    fromTinMSum,
		otherCardsPercent: otherCardPer / 100,
		otherCardsMinSum:  otherCardMSum,
	}
}

var (
	ErrMoneyOnCardOfSenderDontEnough = errors.New("На карте отправителя баланс меньше суммы перевода.")
	ErrTooLowSumOfTransfer = errors.New("Слишком маленькая сумма перевода.")
	ErrInvalidCardNumber = errors.New("Введены неверные данные карты.")
)

func isValid(num string) error {
	num = strings.ReplaceAll(num, " ", "")
	controlSum := int64(0)
	sliceOfNums := strings.Split(num, "")
	for i := 0; i < len(sliceOfNums); i++ {
		value, err := strconv.ParseInt(sliceOfNums[i], 10, 64)
		if err != nil {
			return ErrInvalidCardNumber
		}
		if (i + 1) % 2 != 0 {
			value *= 2
			if value > 9 {
				value -= 9
			}
		}
		controlSum += value
	}

	if controlSum % 10 == 0 {
		return nil
	}

	return ErrInvalidCardNumber
}

func (s *Service) Card2Card(from, to string, amount int64, MCC string, time time.Time) (total int64, err error) {
	errOfValidCardFrom := isValid(from)
	if errOfValidCardFrom != nil {
		fmt.Println("Введены некоректные данные карты.")
		return 0, errOfValidCardFrom
	}

	errOfValidCardTo := isValid(from)
	if errOfValidCardTo != nil {
		fmt.Println("Введены некоректные данные карты.")
		return 0, errOfValidCardFrom
	}

	amountInCents := amount * 100
	errOfFrom, indexOfFrom := s.CardSvc.SearchCards(from)
	errOfTo, indexOfTo := s.CardSvc.SearchCards(to)
	if errOfFrom != nil {
		fmt.Println("Карты с которой вы хотите выполнить перевод нет в нашей базе данных.")
		return 0, card.ErrCardNotInOurBase
	}
	if errOfTo != nil {
		fmt.Println("Карты с которой вы хотите выполнить перевод нет в нашей базе данных.")
		return 0, card.ErrCardNotInOurBase
	}

	if s.CardSvc.StoreOfCards[indexOfFrom].Balance > amountInCents { // Проверяем хватает ли денег на балансе
		if amountInCents > s.fromTinkMinSum {
			s.addTransaction(indexOfFrom, amount, MCC, time)
			s.CardSvc.StoreOfCards[indexOfFrom].Balance -= amountInCents
			s.CardSvc.StoreOfCards[indexOfTo].Balance += amountInCents
			return amount, nil
		} else {
			fmt.Println("Слишком маленькая сумма перевода, введите сумму более 10 руб!")
			return 0, ErrTooLowSumOfTransfer
		}
	}
	fmt.Println("Недостаточно средств на балансе вашей карты.")
	return 0, ErrMoneyOnCardOfSenderDontEnough
}

func (s *Service) addTransaction(index int, amount int64, MCC string, time time.Time) {
		s.CardSvc.StoreOfCards[index].Transactions = append(s.CardSvc.StoreOfCards[index].Transactions, &transaction.Transaction{
		Id:     20,
		Amount: amount * 100,
		MCC:    MCC,
		Date:   time,
		Status: "Completed",
	})
}

func (s *Service) Purchase(amount int64, index int, MCC string, time time.Time) {
	s.addTransaction(index, amount, MCC, time)
	fmt.Println("Сумма вашей покупки ", amount, " рублей")
}

