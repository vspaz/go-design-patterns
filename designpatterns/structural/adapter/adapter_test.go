package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	soapClient := &Soap{}
	restClient := &Rest{}
	soapAdapter := &SoapAdapter{
		Soap: soapClient,
	}

	assert.Equal(t, "making REST request", Get(restClient))
	assert.Equal(t, "making SOAP requests", Get(soapAdapter))
}
