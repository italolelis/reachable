package reachable

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReachable(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario string
		domain   string
		timeout  time.Duration
		function func(*testing.T, string, time.Duration)
	}{
		{
			scenario: "test success lookup domain",
			domain:   "https://google.com",
			timeout:  10 * time.Second,
			function: testLookupDomain,
		},
		{
			scenario: "test success lookup domain with no scheme",
			domain:   "google.com",
			timeout:  10 * time.Second,
			function: testLookupDomain,
		},
		{
			scenario: "test lookup domain and port but no scheme",
			domain:   "google.com:443",
			timeout:  10 * time.Second,
			function: testInvalidDomain,
		},
		{
			scenario: "test invalid domain",
			domain:   "google...wrong.com",
			timeout:  10 * time.Second,
			function: testInvalidDomain,
		},
		{
			scenario: "test success lookup domain with no scheme",
			domain:   "wrongurlfortest.com",
			timeout:  10 * time.Second,
			function: testInvalidDomain,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.function(t, test.domain, test.timeout)
		})
	}
}

func testLookupDomain(t *testing.T, domain string, timeout time.Duration) {
	result, err := IsReachable(context.Background(), domain, timeout)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, result.StatusCode)
}

func testInvalidDomain(t *testing.T, domain string, timeout time.Duration) {
	_, err := IsReachable(context.Background(), domain, timeout)
	assert.Error(t, err)
}
