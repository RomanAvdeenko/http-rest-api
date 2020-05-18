package teststore_test

import (
	"testing"

	"github.com/RomanAvdeenko/http-rest-api/internal/app/model"
	"github.com/RomanAvdeenko/http-rest-api/internal/app/store"
	"github.com/RomanAvdeenko/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "user@example.org"

	_, err := s.User().FindByEmail(email)

	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email

	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	t.Log(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()
	u1 := model.TestUser(t)
	s.User().Create(u1)

	u2, err := s.User().Find(u1.ID)

	assert.NoError(t, err)
	assert.NotNil(t, u2)
	assert.Equal(t, u1.ID, u2.ID)
	assert.Equal(t, u1.Email, u2.Email)
	assert.Equal(t, u1.EncryptedPassword, u2.EncryptedPassword)
}
