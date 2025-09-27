package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Arthut Nicolucci", "arthur@email.com", "1234")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user, user.ID)
	assert.Equal(t, "Arthur Nicolucci", user.Name)
	assert.Equal(t, "arthur@email.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Arthut Nicolucci", "arthur@email.com", "1234")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("1234"))
	assert.False(t, user.ValidatePassword("12345"))
	assert.NotEqual(t, "1234", user.Password)
}
