package whois_test

import (
	"testing"

	"github.com/ak9024/whois/internal/whois"
	"github.com/stretchr/testify/assert"
)

func TestWhois(t *testing.T) {
	wn := whois.New("malascoding.com")
	result, errWhois := wn.Whois()
	assert.Nil(t, errWhois)
	assert.NotNil(t, result)
}
