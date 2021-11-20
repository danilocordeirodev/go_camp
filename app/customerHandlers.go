package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/danilocordeiro/banking/service"
	"github.com/gorilla/mux"
)

type Customer struct {
	Name string `json:"full_name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zipcpde string `json:"zip_code" xml:"zipcode"`
}


type CustomerHandlers struct {
	service service.CustomerService
}



func (ch *CustomerHandlers) getAllCustomers(rw http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomer(status)

	if err != nil {
		writeResponse(rw, err.Code, err.AsMessage(), r)
	} else {
		writeResponse(rw, http.StatusOK, customers, r)
	}
}

func (ch *CustomerHandlers) getCustomer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)

	if err != nil {
		writeResponse(rw, err.Code, err.AsMessage(), r)
	} else {
		writeResponse(rw, http.StatusOK, customer, r)
	}

}

func writeResponse(rw http.ResponseWriter, code int, data interface{}, r *http.Request) {
	if r.Header.Get("Content-Type") == "application/xml" {
		writeResponseXml(rw, code, data)
	} else {
		writeResponseJson(rw, code, data)
	}
}

func writeResponseJson(rw http.ResponseWriter, code int, data interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(code)
	if err := json.NewEncoder(rw).Encode(data); err != nil {
		panic(err)
	}
}

func writeResponseXml(rw http.ResponseWriter, code int, data interface{}) {
	rw.Header().Add("Content-Type", "application/xml")
	rw.WriteHeader(code)
	if err := xml.NewEncoder(rw).Encode(data); err != nil {
		panic(err)
	}

}