package service

import (
	"context"
	"net/http"
)

// AccrualClient определяет интерфейс для обращения к внешней системе начислений (accrual service)

type HTTPAccrualClient struct {
	Client *http.Client
}

func (c *HTTPAccrualClient) GetOrder(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
