package stresstest

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TestCmdSuite struct {
	suite.Suite
}

func (s *TestCmdSuite) SetupSuite() {}

func TestTestCmdSuite(t *testing.T) {
	suite.Run(t, new(TestCmdSuite))
}

func (s *TestCmdSuite) TestHTTPRequests() {
	// Create a mock HTTP server for testing
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set a 200 status code and a JSON response body for testing
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()

	// Set URL to point to our mock server
	url := server.URL + "/test-endpoint"

	// Define the number of requests and concurrency for testing
	requests := 100
	concurrency := 5
	uc := NewStressTest(StressTestDTOInput{
		Url:         url,
		Requests:    requests,
		Concurrency: concurrency,
	})

	res, err := uc.Execute()

	s.NoError(err)
	requestsMade := 0
	for _, tot := range res.Results {
		requestsMade += tot
	}
	s.Equal(requestsMade, requests)
	s.Greater(res.ExecutionTime, time.Duration(0))

}
