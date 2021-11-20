package app

import (
	"encoding/json"
	"net/http"

	"github.com/danilocordeiro/banking/dto"
	"github.com/danilocordeiro/banking/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}


func (ah AccountHandler) NewAccount(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(rw, http.StatusBadRequest, err.Error(), r)
	} else {
		request.CustomerId = customerId
		account, appError := ah.service.NewAccount(request)
		if appError != nil {
			writeResponse(rw, appError.Code, appError.Message, r)
		} else {
			writeResponse(rw, http.StatusCreated, account, r)
		}
	}

}

func (ah AccountHandler) MakeTransaction(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	accountId := vars["account_id"]
	var request dto.TransactionRequest
	
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(rw, http.StatusBadRequest, err.Error(), r)
	} else {
		request.CustomerId = customerId
		request.AccountId  = accountId

		transaction, appError := ah.service.MakeTransaction(request)
		if appError != nil {
			writeResponse(rw, appError.Code, appError.Message, r)
		} else {
			writeResponse(rw, http.StatusCreated, transaction, r)
		}
	}

}
