package whois

import (
	_whois "github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

type whois struct {
	Domain string
}

type IWhois interface {
	Whois() (*whoisparser.WhoisInfo, error)
}

func New(domain string) IWhois {
	return &whois{
		Domain: domain,
	}
}

func (w *whois) Whois() (*whoisparser.WhoisInfo, error) {
	raw, errWhois := _whois.Whois(w.Domain)
	if errWhois != nil {
		return nil, errWhois
	}

	result, errWhoisParse := whoisparser.Parse(raw)
	if errWhoisParse != nil {
		return nil, errWhoisParse
	}

	return &result, nil
}
