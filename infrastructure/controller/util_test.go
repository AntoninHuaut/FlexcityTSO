package controller

import (
	"FlexcityTest/domain"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUtil_IsFutureDate(t *testing.T) {
	futureTime := time.Now().Add(time.Hour)
	pastTime := time.Now().Add(-time.Hour)

	type testStruct struct {
		Date *time.Time `validate:"is-future-date"`
	}

	tests := []struct {
		name   string
		args   testStruct
		checks func(t *testing.T, err error)
	}{
		{
			name: "valid future date",
			args: testStruct{Date: &futureTime},
			checks: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "invalid past date",
			args: testStruct{Date: &pastTime},
			checks: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "nil date",
			args: testStruct{Date: nil},
			checks: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Struct(tt.args)
			tt.checks(t, err)
		})
	}
}

func TestUtil_ValidatePayload(t *testing.T) {
	type testStruct struct {
		Number *int `validate:"required,min=2"`
	}
	int5 := 5
	int1 := 1

	tests := []struct {
		name   string
		args   testStruct
		checks func(t *testing.T, err error)
	}{
		{
			name: "valid payload",
			args: testStruct{Number: &int5},
			checks: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "nil payload",
			args: testStruct{Number: nil},
			checks: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "invalid payload",
			args: testStruct{Number: &int1},
			checks: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validatePayload(tt.args)
			tt.checks(t, err)
		})
	}
}

func TestUtil_HandleError(t *testing.T) {
	type testStruct struct {
		Err error
	}

	tests := []struct {
		name   string
		args   testStruct
		checks func(t *testing.T, code int)
	}{
		{
			name: "nil error",
			args: testStruct{Err: nil},
			checks: func(t *testing.T, code int) {
				assert.Equal(t, http.StatusOK, code)
			},
		},
		{
			name: "internal error",
			args: testStruct{Err: errors.New("internal failure")},
			checks: func(t *testing.T, code int) {
				assert.Equal(t, http.StatusInternalServerError, code)
			},
		},
		{
			name: "domain error: invalid payload",
			args: testStruct{Err: domain.ErrorResponse{
				Type:        domain.ErrInvalidPayload,
				NativeError: errors.New("invalid payload"),
			}},
			checks: func(t *testing.T, code int) {
				assert.Equal(t, http.StatusBadRequest, code)
			},
		},
		{
			name: "domain error: internal",
			args: testStruct{Err: domain.ErrorResponse{
				Type:        domain.ErrInternal,
				NativeError: errors.New("internal failure"),
			}},
			checks: func(t *testing.T, code int) {
				assert.Equal(t, http.StatusInternalServerError, code)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			handleError(w, req, tt.args.Err)
			tt.checks(t, w.Code)
		})
	}
}
