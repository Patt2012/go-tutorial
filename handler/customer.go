package handler

import (
	"bank/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerHandler struct {
	custService service.CustomerService
}

func NewCustomerHandler(custService service.CustomerService) customerHandler {
	return customerHandler{custService: custService}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	cusomers, err := h.custService.GetCustomers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(cusomers)
}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {

	customerID, _ :=strconv.Atoi(mux.Vars(r)["customerID"])

	customer, err := h.custService.GetCustomer(customerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customer)
}