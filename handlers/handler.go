package handlers

import (
	"fmt"
	"gymshark/service"
	"net/http"
	"strconv"
	"strings"
)

// Amount handler accepts amount as a form parameter and returns result map[pack1:amount1 pack2:amount2 ...]
func Amount(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	strAmount := r.Form.Get("amount")
	amount, err := strconv.Atoi(strAmount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad amount parameter"))
		return
	}

	packs := service.PackNumber(amount)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, packs)
}

// Packs handler accept packs array as a form parameter and returns 200 if no errors
func Packs(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	strPacks := r.Form.Get("packs")
	s := strings.Split(strPacks, ",")

	if len(s) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad packs parameter"))
		return
	}

	err := service.Packs(s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad packs parameter"))
		return
	}

	w.WriteHeader(http.StatusOK)
}
