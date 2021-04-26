package reachable

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/italolelis/reachable/pkg/log"
	httpstat "github.com/tcnksm/go-httpstat"
)

type Target struct {
	URL    string
	Method string
}

func NewTarget(url string) *Target {
	return &Target{url, http.MethodGet}
}

// Reachable represents the response for a domain check
type Reachable struct {
	Domain     string
	Port       string
	IP         string
	StatusCode int
	Response   *httpstat.Result
}

// IsReachable checks if a domain is reachable
func IsReachable(ctx context.Context, t *Target) (Reachable, error) {
	u, err := url.Parse(t.URL)
	if err != nil {
		return Reachable{}, err
	}

	if u.Scheme == "" {
		u.Scheme = "http"
	}

	var result httpstat.Result

	ctx = httpstat.WithHTTPStat(ctx, &result)
	req, err := http.NewRequestWithContext(ctx, t.Method, u.String(), nil)
	if err != nil {
		return Reachable{}, err
	}

	r := Reachable{
		Domain:   u.Hostname(),
		Port:     u.Port(),
		Response: &result,
	}

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return r, err
	}

	r.StatusCode = res.StatusCode

	if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
		return r, err
	}

	res.Body.Close()
	result.End(time.Now())

	return r, nil
}

// IsReachableAsync checks if a domain is reachable in a async way
func IsReachableAsync(ctx context.Context, ch chan<- Reachable, chErr chan<- error, urls ...string) {
	logger := log.WithContext(ctx)

	var wg sync.WaitGroup
	for _, t := range urls {
		wg.Add(1)
		go func(url string) {
			logger.Debugw("checking domain", "domain", url)
			result, err := IsReachable(ctx, NewTarget(url))
			if err != nil {
				chErr <- fmt.Errorf("unreachable: %w", err)
			}
			logger.Debugw("Reachable!", "domain", url)

			ch <- result

			wg.Done()
		}(t)
	}

	wg.Wait()
	close(ch)
}
