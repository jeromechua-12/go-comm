package auth

import (
	"maps"
	"testing"
)

func TestValidateSignup(t *testing.T) {
	tests := []struct {
		name string
		input UserSignupForm
		want map[string]string
	}{
		{
			name: "Valid",
			input: UserSignupForm{
				Name: "Alex",
				Email: "alex123@gmail.com",
				Password: "password123",
			},
			want: make(map[string]string),
		},
		{
			name: "EmptyFields",
			input: UserSignupForm{
				Name: "",
				Email: "",
				Password: "",
			},
			want: map[string]string{
				"name": "Name cannot be empty",
				"email": "Email cannot be empty",
				"password": "Password cannot be empty",
			},
		},
		{
			name: "InvalidEmail",
			input: UserSignupForm{
				Name: "Bob",
				Email: "bobgmail.com",
				Password: "password123",
			},
			want: map[string]string{
				"email": "Invalid email address",
			},
		},
		{
			name: "ShortPassword",
			input: UserSignupForm{
				Name: "Carmelo",
				Email: "CarmeloAnthony@yahoo.com",
				Password: "pwd",
			},
			want: map[string]string{
				"password": "Password must be at least 8 characters",
			},
		},
		{
			name: "LongPassword",
			input: UserSignupForm{
				Name: "Carmelo",
				Email: "CarmeloAnthony@yahoo.com",
				Password: "ééééééééééééééééééééééééééééééééééééé",
			},
			want: map[string]string{
				"password": "Password cannot be longer than 72 bytes",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fieldErrors := validateSignup(tt.input)
			if !maps.Equal(fieldErrors, tt.want) {
				t.Errorf("got %v; want %v", fieldErrors, tt.want)
			}
		})
	}
}
