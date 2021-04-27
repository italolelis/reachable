package log

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario string
		function func(*testing.T)
	}{
		{
			scenario: "create log context",
			function: testCreateLogContext,
		},
		{
			scenario: "create log context with nil context",
			function: testCreateLogWithNilContext,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.function(t)
		})
	}
}

func testCreateLogContext(t *testing.T) {
	ctx := NewContext(context.Background())
	assert.NotNil(t, ctx)

	logger := WithContext(ctx)
	assert.NotNil(t, logger)
}

func testCreateLogWithNilContext(t *testing.T) {
	assert.Panics(t, func() { NewContext(nil) })

	logger := WithContext(nil)
	assert.NotNil(t, logger)
}
