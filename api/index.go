package api

import (
	"encoding/json"
	"net/http"

	"github.com/likexian/whois"
)

type WhoisResponse20xOk struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
}

type WhoisRequest struct {
	Domain string `json:"domain"`
}

func Handler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "application/json")

	var whoisRequest WhoisRequest
	errDecode := json.NewDecoder(r.Body).Decode(&whoisRequest)
	if errDecode != nil {
		return errDecode
	}

	result, errWhois := whois.Whois(whoisRequest.Domain)
	if errWhois != nil {
		return errWhois
	}

	errEncode := json.NewEncoder(w).Encode(WhoisResponse20xOk{
		Code:   http.StatusOK,
		Result: result,
	})
	if errEncode != nil {
		return errEncode
	}

	return nil
}
