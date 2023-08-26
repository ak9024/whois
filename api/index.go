package api

import (
	"encoding/json"
	"net/http"

	"github.com/ak9024/whois/whois"
)

var whoisRequest whois.WhoisRequest

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
	}

	errDecode := json.NewDecoder(r.Body).Decode(&whoisRequest)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	wn := whois.New(whoisRequest.Domain)
	result, errWhois := wn.Whois()
	if errWhois != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	errEncode := json.NewEncoder(w).Encode(whois.WhoisResponse20xOk{
		Code:   http.StatusOK,
		Result: result,
	})
	if errEncode != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
