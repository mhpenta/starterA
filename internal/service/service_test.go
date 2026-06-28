package service

import (
	"errors"
	"testing"
)

func TestValidateCreateUserInputTrimsValidFields(t *testing.T) {
	input := &CreateUserInput{
		Username: "  marc  ",
		Email:    "  marc@example.com  ",
	}

	err := validateCreateUserInput(input)
	if err != nil {
		t.Fatalf("validate create input: %v", err)
	}
	if input.Username != "marc" {
		t.Fatalf("Username = %q, want marc", input.Username)
	}
	if input.Email != "marc@example.com" {
		t.Fatalf("Email = %q, want marc@example.com", input.Email)
	}
}

func TestValidateCreateUserInputRejectsInvalidFields(t *testing.T) {
	tests := []struct {
		name  string
		input *CreateUserInput
	}{
		{name: "nil", input: nil},
		{name: "missing username", input: &CreateUserInput{Email: "marc@example.com"}},
		{name: "short username", input: &CreateUserInput{Username: "ma", Email: "marc@example.com"}},
		{name: "missing email", input: &CreateUserInput{Username: "marc"}},
		{name: "invalid email", input: &CreateUserInput{Username: "marc", Email: "not-an-email"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCreateUserInput(tt.input)
			if !errors.Is(err, ErrInvalidUserInput) {
				t.Fatalf("err = %v, want %v", err, ErrInvalidUserInput)
			}
		})
	}
}

func TestValidateUpdateUserInputRejectsInvalidFields(t *testing.T) {
	input := &UpdateUserInput{
		Username: "marc",
		Email:    "not-an-email",
	}

	err := validateUpdateUserInput(input)
	if !errors.Is(err, ErrInvalidUserInput) {
		t.Fatalf("err = %v, want %v", err, ErrInvalidUserInput)
	}
}
