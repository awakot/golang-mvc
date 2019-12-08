package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "userid 0 is not existed")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNotError(t *testing.T) {
	user, err := GetUser(0001)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 0001, user.ID)
	assert.EqualValues(t, "hoge", user.FirstNname)
	assert.EqualValues(t, "murakawa", user.LastName)
	assert.EqualValues(t, "hoge@example.com", user.Email)

}
