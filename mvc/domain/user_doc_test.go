package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "we were not expecting a user")
	assert.NotNil(t, err, "we were expecting an error when user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode, "we were expecting 404 when user not found")
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
}
