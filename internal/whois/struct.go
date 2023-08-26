package whois

type WhoisResponse20xOk struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
}

type WhoisRequest struct {
	Domain string `json:"domain"`
}
