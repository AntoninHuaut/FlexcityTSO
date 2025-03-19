package domain

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestErrorResponse(t *testing.T) {
	tests := []struct {
		name        string
		m           ErrorResponse
		wantMessage string
		wantCode    int
	}{
		{
			name: "database error response",
			m: ErrorResponse{
				NativeError: nil,
				Type:        ErrDatabase,
			},
			wantMessage: "database",
			wantCode:    500,
		},
		{
			name: "invalid payload error response",
			m: ErrorResponse{
				NativeError: nil,
				Type:        ErrInvalidPayload,
			},
			wantMessage: "invalid_payload",
			wantCode:    400,
		},
		{
			name: "unknown error response",
			m: ErrorResponse{
				NativeError: nil,
				Type:        "unknown_test",
			},
			wantMessage: "unknown_test",
			wantCode:    500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.wantMessage, tt.m.Error())
			require.Equal(t, tt.wantCode, tt.m.StatusCode())
		})
	}
}
