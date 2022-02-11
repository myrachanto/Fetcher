package fetcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	fetcher := CreateFetcher("")
	err := fetcher.Validate()
	expected := "Invalid EndPoint"
	assert.EqualValues(t, err.Message(), expected, "the url must not be zero")
}
