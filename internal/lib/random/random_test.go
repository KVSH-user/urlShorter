package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRandomString(t *testing.T) {
	tests := []struct {
		name        string
		aliasLength int
	}{
		{
			name:        "aliasLength = 1",
			aliasLength: 1,
		},
		{
			name:        "aliasLength = 5",
			aliasLength: 5,
		},
		{
			name:        "aliasLength = 10",
			aliasLength: 10,
		},
		{
			name:        "aliasLength = 20",
			aliasLength: 20,
		},
		{
			name:        "aliasLength = 30",
			aliasLength: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str1 := NewRandomString(tt.aliasLength)
			str2 := NewRandomString(tt.aliasLength)

			assert.Len(t, str1, tt.aliasLength)
			assert.Len(t, str2, tt.aliasLength)

			// Check that two generated strings are different
			// This is not an absolute guarantee that the function works correctly,
			// but this is a good heuristic for a simple random generator.
			assert.NotEqual(t, str1, str2)
		})
	}
}
