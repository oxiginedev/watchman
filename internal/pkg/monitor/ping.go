package monitor

import (
	"context"
	"errors"
	"net/http"
	"net/url"
)

var ErrInvalidURL = errors.New("invalid url format")

type PingResponse struct {
	Up     bool
	Status int
}

func Ping(ctx context.Context, s string) (*PingResponse, error) {
	parsedURL, err := url.Parse(s)
	if err != nil {
		return nil, ErrInvalidURL
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &PingResponse{Up: false, Status: res.StatusCode}, nil
	}

	defer res.Body.Close()

	up := res.StatusCode < http.StatusBadRequest

	return &PingResponse{Up: up, Status: res.StatusCode}, nil
}
