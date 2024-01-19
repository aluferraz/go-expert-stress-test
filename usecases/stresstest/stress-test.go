package stresstest

import (
	"errors"
	"net/http"
	"sync"
	"time"
)

type StressTestDTOInput struct {
	Url         string
	Requests    int
	Concurrency int
}

type StressTestDTOOutput struct {
	Results       map[int]int
	ExecutionTime time.Duration
}

type StressTest struct {
	Input StressTestDTOInput
}

func NewStressTest(input StressTestDTOInput) *StressTest {

	return &StressTest{
		input,
	}
}

func (st *StressTest) Execute() (StressTestDTOOutput, error) {
	start := time.Now()

	var wg sync.WaitGroup
	queue := make(chan struct{}, st.Input.Concurrency)
	defer close(queue)
	requests := st.Input.Requests
	resultsChan := make(chan int, requests)
	defer close(resultsChan)
	result := StressTestDTOOutput{
		Results: make(map[int]int),
	}

	var hasError error
	for i := 0; i < requests; i++ {
		wg.Add(1)
		queue <- struct{}{}

		go func() {
			resp, err := http.Get(st.Input.Url)
			if err != nil {
				hasError = errors.Join(hasError, err)
				<-queue
				return
			}
			resultsChan <- resp.StatusCode
			<-queue
		}()
	}
	go func() {
		for res := range resultsChan {
			result.Results[res] += 1
			wg.Done()
		}
	}()

	wg.Wait()
	result.ExecutionTime = time.Since(start)
	return result, hasError
}
