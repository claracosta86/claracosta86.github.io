package notification_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"poc2/back/domain/notification"
)

func TestNewNotificationType(t *testing.T) {
	tests := []struct {
		name        string
		inputType   string
		expected    notification.NotificationType
		expectError bool
	}{
		{
			name:        "valid commented",
			inputType:   "commented",
			expected:    notification.NotificationTypeCommented,
			expectError: false,
		},
		{
			name:        "valid canceled",
			inputType:   "canceled",
			expected:    notification.NotificationTypeCanceled,
			expectError: false,
		},
		{
			name:        "valid updated",
			inputType:   "updated",
			expected:    notification.NotificationTypeUpdated,
			expectError: false,
		},
		{
			name:        "invalid type",
			inputType:   "invalid",
			expected:    "",
			expectError: true,
		},
		{
			name:        "empty type",
			inputType:   "",
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nt, err := notification.NewNotificationType(tt.inputType)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, nt)
				assert.Equal(t, tt.inputType, nt.String())
			}
		})
	}
}
