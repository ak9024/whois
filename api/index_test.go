package api_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ak9024/whois/api"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	buf := []byte(`{"domain": "malascoding.com"}`)
	req := httptest.NewRequest(http.MethodPost, "/api", bytes.NewBuffer(buf))
	w := httptest.NewRecorder()
	errHandler := api.Handler(w, req)
	assert.Nil(t, errHandler)
	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)
	body, errBody := io.ReadAll(resp.Body)
	assert.Nil(t, errBody)
	t.Log(string(body))
}
