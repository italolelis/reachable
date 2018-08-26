package reachable

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gojektech/heimdall/httpclient"
	httpstat "github.com/tcnksm/go-httpstat"
)

// Reachable represents the response for a domain check
type Reachable struct {
	Domain     string
	Port       string
	IP         string
	StatusCode int
	Response   httpstat.Result
}

// IsReachable checks if a domain is reachable
func IsReachable(ctx context.Context, domain string, timeout time.Duration) (*Reachable, error) {
	var result httpstat.Result

	u, err := url.Parse(domain)
	if err != nil {
		return nil, err
	}

	if u.Scheme == "" {
		u.Scheme = "http"
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	ctx = httpstat.WithHTTPStat(ctx, &result)
	req = req.WithContext(ctx)

	c := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
		return nil, err
	}
	res.Body.Close()
	result.End(time.Now())

	r := Reachable{
		Domain:     res.Request.URL.Hostname(),
		Port:       res.Request.URL.Port(),
		StatusCode: res.StatusCode,
		Response:   result,
	}

	return &r, nil
}
