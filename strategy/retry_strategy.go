package strategy

import (
	"github.com/avast/retry-go/v4"
	"log"
	"net/http"
	"time"
)

func SendRequest(httpReq string, attempt int8) (*http.Response, error) {
	var resp *http.Response

	configs := []retry.Option{
		retry.Attempts(uint(attempt)),
		retry.OnRetry(func(n uint, err error) {
			log.Printf("Retry request %d to and get error: %v", n+1, err)
		}),
		retry.Delay(time.Second),
	}

	err := retry.Do(
		func() error {
			var err error
			resp, err = http.Get(httpReq)
			return err
		},
		configs...,
	)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return resp, nil
}
