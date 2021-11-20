package dto

import (
	"strings"

	"github.com/danilocordeiro/banking/errs"
)


const WITHDRAWAL = "withdrawal"
const DEPOSIT    = "deposit"

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	CustomerId      string  `json:"-"`
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"new_balance"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (t TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return t.TransactionType == WITHDRAWAL
}

func (t TransactionRequest) Validate() *errs.AppError {
	if  strings.ToLower(t.TransactionType) != WITHDRAWAL  &&  strings.ToLower(t.TransactionType) != DEPOSIT{
		return errs.NewValidationError("Transition type can only be deposit or withdrawal")
	}

	if t.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}

	return nil
}