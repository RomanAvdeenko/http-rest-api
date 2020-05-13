package model_test

import (
	"testing"

	"github.com/RomanAvdeenko/http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "Valid email",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "Empty email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "Invalid email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalid@"
				return u
			},
			isValid: false,
		},
		{
			name: "Valid password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "123456"
				return u
			},
			isValid: true,
		},
		{
			name: "With encrypted password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encrypted_password"
				return u
			},
			isValid: true,
		},
		{
			name: "Empty password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "Short password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "911"
				return u
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotNil(t, u.EncryptedPassword)
}
