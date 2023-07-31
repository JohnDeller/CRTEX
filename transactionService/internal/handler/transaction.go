package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"strings"
	"time"
	"transactionService/pkg/proto/transaction_v1"
)

const (
	operationTypeCrediting = "crediting"
	operatioinTypeDebiting = "debiting"

	currencyUsd = "USD"
	currencyEur = "EUR"
	currencyRub = "RUB"
)

func (h *Handler) createTransaction(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Println("panic in createTransaction", err)
		}
	}()
	var transaction Transaction
	if err := c.BindJSON(&transaction); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	transaction.CreatedAt = time.Now()
	transactionDto := newTransactionToDto(transaction)
	// todo add support more currency
	if transactionDto.GetCurrency() != transaction_v1.Currency_CURRENCY_USD {
		newErrorResponse(c, http.StatusBadRequest, "incorrect currency, need USD")
		return
	}
	if transactionDto.GetOperationType() == transaction_v1.OperationType_OPERATION_TYPE_NONE {
		newErrorResponse(c, http.StatusBadRequest, "incorrect operation type, need crediting/debiting")
		return
	}
	id, err := h.services.TransactionsService.Create(transactionDto)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	transaction.Id = id
	c.JSON(http.StatusOK, transaction)
}

func (h *Handler) getTransactionById(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Println("panic in getTransactionById", err)
		}
	}()
	id := c.Param("id")
	transaction, err := h.services.TransactionsService.GetById(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newTransactionFromDto(transaction))
}

func operationTypeToResponse(operationType transaction_v1.OperationType) string {
	switch operationType {
	case transaction_v1.OperationType_OPERATION_TYPE_CREDITING:
		return operationTypeCrediting
	case transaction_v1.OperationType_OPERATION_TYPE_DEBITING:
		return operatioinTypeDebiting

	default:
		return ""
	}
}

func operationTypeFromRequest(operationType string) transaction_v1.OperationType {
	switch strings.ToLower(operationType) {
	case operationTypeCrediting:
		return transaction_v1.OperationType_OPERATION_TYPE_CREDITING
	case operatioinTypeDebiting:
		return transaction_v1.OperationType_OPERATION_TYPE_DEBITING

	default:
		return transaction_v1.OperationType_OPERATION_TYPE_NONE
	}
}

func currencyToResponse(currency transaction_v1.Currency) string {
	switch currency {
	case transaction_v1.Currency_CURRENCY_USD:
		return currencyUsd
	case transaction_v1.Currency_CURRENCY_EUR:
		return currencyEur
	case transaction_v1.Currency_CURRENCY_RUB:
		return currencyRub

	default:
		return ""
	}
}

func currencyFromRequest(currency string) transaction_v1.Currency {
	switch strings.ToUpper(currency) {
	case currencyUsd:
		return transaction_v1.Currency_CURRENCY_USD
	case currencyEur:
		return transaction_v1.Currency_CURRENCY_EUR
	case currencyRub:
		return transaction_v1.Currency_CURRENCY_RUB

	default:
		return transaction_v1.Currency_CURRENCY_NONE
	}
}

type Transaction struct {
	Id            string    `json:"id,omitempty"`
	UserId        string    `json:"user_id,omitempty"`
	OperationType string    `json:"operation_type,omitempty"`
	Price         float64   `json:"price,omitempty"`
	Currency      string    `json:"currency,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}

func newTransactionToDto(transaction Transaction) transaction_v1.Transaction {
	return transaction_v1.Transaction{
		Id:            transaction.Id,
		UserId:        transaction.UserId,
		OperationType: operationTypeFromRequest(transaction.OperationType),
		Price:         transaction.Price,
		Currency:      currencyFromRequest(transaction.Currency),
		CreatedTime:   timestamppb.New(transaction.CreatedAt),
	}
}

func newTransactionFromDto(transaction transaction_v1.Transaction) Transaction {
	return Transaction{
		Id:            transaction.GetId(),
		UserId:        transaction.GetUserId(),
		OperationType: operationTypeToResponse(transaction.GetOperationType()),
		Price:         transaction.GetPrice(),
		Currency:      currencyToResponse(transaction.GetCurrency()),
		CreatedAt:     transaction.CreatedTime.AsTime(),
	}
}
