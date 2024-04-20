package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type EndPointServices struct {
	listenAddr string
	store      *PgClient
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ServiceError struct {
	ErrorMsg string
}

func NewEndPoint(listenAddr string, store *PgClient) *EndPointServices {
	return &EndPointServices{
		listenAddr: listenAddr,
		store:      store,
	}

}

func AttachJSON(w http.ResponseWriter, status int, msg any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(msg)
}
func makeHttpHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle the error
			AttachJSON(w, http.StatusBadRequest, ServiceError{ErrorMsg: err.Error()})
		}

	}
}

func (s *EndPointServices) Fire() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHttpHandler(s.handleEntity))
	router.HandleFunc("/account/{id}", makeHttpHandler(s.handleGetEntityById))

	log.Println("THe api is running")

	http.ListenAndServe(s.listenAddr, router)

}
func (s *EndPointServices) handleEntity(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		return s.handleGetEntity(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateEntity(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteEntity(w, r)
	}
	fmt.Errorf("Method is not allowed")
	return nil

}
func (s *EndPointServices) handleCreateEntity(w http.ResponseWriter, r *http.Request) error {
	createAccReq := new(BodyReq)
	if err := json.NewDecoder(r.Body).Decode(createAccReq); err != nil {
		return err
	}

	account := NewAccount(createAccReq.FirstName, createAccReq.LastName)
	if err := s.store.CreateAccount(account); err != nil {
		return nil
	}
	return AttachJSON(w, http.StatusOK, account)

}

func (s *EndPointServices) handleGetEntity(w http.ResponseWriter, r *http.Request) error {

	accounts, err := s.store.SelectAllAccount()
	if err != nil {
		return err
	}
	return AttachJSON(w, http.StatusOK, accounts)
}

func (s *EndPointServices) handleGetEntityById(w http.ResponseWriter, r *http.Request) error {
	id, err := parseIdParas(r)
	if err != nil {
		return err
	}
	if r.Method == "GET" {

		accounts, err := s.store.GetAccountById(id)
		if err != nil {
			return err
		}
		return AttachJSON(w, http.StatusOK, accounts)
	}
	if r.Method == "DELETE" {
		err := s.store.DeleteAccount(id)
		if err != nil {
			return nil
		}

		return AttachJSON(w, http.StatusOK, map[string]int{"id": id})

	}

	return fmt.Errorf("Method not allowed")
}
func (s *EndPointServices) handleDeleteEntity(w http.ResponseWriter, r *http.Request) error {
	id, err := parseIdParas(r)
	if err != nil {
		return err

	}
	if err := s.store.DeleteAccount(id); err != nil {
		return nil
	}
	return AttachJSON(w, http.StatusOK, map[string]int{"id": id})
}

func (s *EndPointServices) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	transferAcc := Transfer{}
	if err := json.NewDecoder(r.Body).Decode(&transferAcc); err != nil {
		return nil
	}
	balance, err := s.store.CheckSenderBalance(transferAcc.formAccount)
	if err != nil {
		return nil
	}
	if balance <  int64(transferAcc.amount) {
		
		return fmt.Errorf("Not enough money")

	}
	return nil

}
func (s *EndPointServices) hanleWithdraw(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func parseIdParas(r *http.Request) (int, error) {
	ids := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(ids)
	if err != nil {
		return -1, fmt.Errorf("Invalid Params")
	}
	return idInt, nil

}
